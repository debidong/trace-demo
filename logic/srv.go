package logic

import (
	"context"
	"net"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

const SrvNameGorm = "gorm"

func NewServer(srvName string) *http.Server {
	ctx := context.Background()
	config := C.Server
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

	uris := C.Server[srvName].Uri
	switch srvName {
	case SrvNameGorm:
		for _, uri := range uris {
			handleFunc(uri, handleGorm)
		}
	default:
		for _, uri := range urls {
			handleFunc(uri, handleCommon)
		}
	}

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
