package country

import (
	"context"
	"net/http"
	"trace-demo/server/otel"
)

func America(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "America")
	r.WithContext(ctx)
	defer span.End()

	americaPool.GetRandomCity()(w, r)
}

// new-york -> los-angeles -> chicago
func NewYork(w http.ResponseWriter, r *http.Request) {
	newyork(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New York"))
}

// los-angeles -> chicago
func LosAngeles(w http.ResponseWriter, r *http.Request) {
	losangeles(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Los Angeles"))
}

func Chicago(w http.ResponseWriter, r *http.Request) {
	chicago(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chicago"))
}

// to england
func Edmond(w http.ResponseWriter, r *http.Request) {
	edmond(r.Context())
	England(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edmond"))
}

func newyork(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "New York")
	defer span.End()

	losangeles(ctx)
}

func losangeles(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Los Angeles")
	defer span.End()

	chicago(ctx)
}

func chicago(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Chicago")
	defer span.End()
}

func edmond(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Edmond")
	defer span.End()
}
