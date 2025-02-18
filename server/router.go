package server

import (
	"net/http"
	"trace-demo/server/country"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func NewHTTPHandler() http.Handler {
	mux := http.NewServeMux()

	// handleFunc is a replacement for mux.HandleFunc
	// which enriches the handler's HTTP instrumentation with the pattern as the http.route.
	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		mux.Handle(pattern, handler)
	}

	// Register handlers.
	handleFunc("/china", country.China)
	handleFunc("/china/beijing", country.Beijing)
	handleFunc("/china/shanghai", country.Shanghai)
	handleFunc("/china/guangzhou", country.Guangzhou)
	handleFunc("/china/handan", country.Handan)

	handleFunc("/america", country.America)
	handleFunc("/america/new-york", country.NewYork)
	handleFunc("/america/los-angeles", country.LosAngeles)
	handleFunc("/america/chicago", country.Chicago)
	handleFunc("/america/edmond", country.Edmond)

	handleFunc("/england", country.England)
	handleFunc("/england/london", country.London)
	handleFunc("/england/manchester", country.Manchester)
	handleFunc("/england/liverpool", country.Liverpool)
	handleFunc("/england/edinburgh", country.Edinburgh)

	// Add HTTP instrumentation for the whole server.
	handler := otelhttp.NewHandler(mux, "/")
	return handler
}
