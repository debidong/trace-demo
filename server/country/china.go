package country

import (
	"net/http"
	"trace-demo/server/otel"
)

func China(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "China Custom")
	r.WithContext(ctx)
	defer span.End()

	chinaPool.GetRandomCity()(w, r)
}

// beijing -> shanghai -> guangzhou
func Beijing(w http.ResponseWriter, r *http.Request) {
	beijing(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Beijing"))
}

// shanghai -> guangzhou
func Shanghai(w http.ResponseWriter, r *http.Request) {
	shanghai(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shanghai"))
}

func Guangzhou(w http.ResponseWriter, r *http.Request) {
	guangzhou(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Guangzhou"))
}

// to america
func Handan(w http.ResponseWriter, r *http.Request) {
	handan(r)
	America(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Handan"))
}

func beijing(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Beijing")
	r.WithContext(ctx)
	defer span.End()

	shanghai(r)
}

func shanghai(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Shanghai")
	r.WithContext(ctx)
	defer span.End()

	guangzhou(r)
}

func guangzhou(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Guangzhou")
	r.WithContext(ctx)
	defer span.End()
}

func handan(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Handan")
	r.WithContext(ctx)
	defer span.End()
}
