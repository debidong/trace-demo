package logic

import (
	"context"
	"math/rand"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func NewServer(srvName string) *http.Server {
	ctx := context.Background()
	config := MustLoadConfig().Server
	// Start HTTP server.
	return &http.Server{
		Addr:         config[srvName].Addr,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      newHTTPHandler(srvName),
	}
}

func newHTTPHandler(srvName string) http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	urls := []string{}
	for srvName, srv := range MustLoadConfig().Server {
		for _, uri := range srv.Uri {
			urls = append(urls, FormatRequestURL(srvName, uri))
		}
	}

	_handler := func(w http.ResponseWriter, r *http.Request) {
		ctx, span := tracer.Start(r.Context(), r.URL.Path)
		defer span.End()

		// whether to return immediately
		if rand.Intn(2) == 0 {
			return
		}

		// get resource
		next := rand.Intn(len(urls))
		req, _ := http.NewRequestWithContext(ctx, "GET", urls[next], nil)

		propagator := otel.GetTextMapPropagator()
		propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

		client := &http.Client{
			Transport: http.DefaultTransport,
		}
		client.Do(req)
		return
	}

	for _, uri := range MustLoadConfig().Server[srvName].Uri {
		handleFunc(uri, _handler)
	}

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
