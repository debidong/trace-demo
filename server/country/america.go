package country

import (
	"net/http"
	"trace-demo/server/otel"
)

func America(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "America Custom")
	r.WithContext(ctx)
	defer span.End()

	americaPool.GetRandomCity()(w, r)
}

// new-york -> los-angeles -> chicago
func NewYork(w http.ResponseWriter, r *http.Request) {
	newyork(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New York"))
}

// los-angeles -> chicago
func LosAngeles(w http.ResponseWriter, r *http.Request) {
	losangeles(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Los Angeles"))
}

func Chicago(w http.ResponseWriter, r *http.Request) {
	chicago(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chicago"))
}

// to england
func Edmond(w http.ResponseWriter, r *http.Request) {
	edmond(r)
	England(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edmond"))
}

func newyork(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "New York")
	r.WithContext(ctx)
	defer span.End()

	losangeles(r)
}

func losangeles(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Los Angeles")
	r.WithContext(ctx)
	defer span.End()

	chicago(r)
}

func chicago(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Chicago")
	r.WithContext(ctx)
	defer span.End()
}

func edmond(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Edmond")
	r.WithContext(ctx)
	defer span.End()
}
