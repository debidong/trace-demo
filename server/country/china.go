package country

import (
	"context"
	"net/http"
	"trace-demo/server/otel"
)

func China(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "China")
	r.WithContext(ctx)
	defer span.End()

	chinaPool.GetRandomCity()(w, r)
}

// beijing -> shanghai -> guangzhou
func Beijing(w http.ResponseWriter, r *http.Request) {
	beijing(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Beijing"))
}

// shanghai -> guangzhou
func Shanghai(w http.ResponseWriter, r *http.Request) {
	shanghai(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shanghai"))
}

func Guangzhou(w http.ResponseWriter, r *http.Request) {
	guangzhou(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Guangzhou"))
}

// to america
func Handan(w http.ResponseWriter, r *http.Request) {
	handan(r.Context())
	America(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Handan"))
}

func beijing(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Beijing")
	defer span.End()

	shanghai(ctx)
}

func shanghai(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Shanghai")
	defer span.End()

	guangzhou(ctx)
}

func guangzhou(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Guangzhou")
	defer span.End()
}

func handan(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Handan")
	defer span.End()
}
