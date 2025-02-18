package country

import (
	"context"
	"net/http"
	"trace-demo/server/otel"
)

func England(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "England")
	r.WithContext(ctx)
	defer span.End()

	englandPool.GetRandomCity()(w, r)
}

// london -> manchester -> liverpool
func London(w http.ResponseWriter, r *http.Request) {
	london(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("London"))
}

// manchester -> liverpool
func Manchester(w http.ResponseWriter, r *http.Request) {
	manchester(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Manchester"))
}

func Liverpool(w http.ResponseWriter, r *http.Request) {
	liverpool(r.Context())
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Liverpool"))
}

// to china
func Edinburgh(w http.ResponseWriter, r *http.Request) {
	edinburgh(r.Context())
	China(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edinburgh"))
}

func london(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "London")
	defer span.End()

	manchester(ctx)
}

func manchester(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Manchester")
	defer span.End()

	liverpool(ctx)
}

func liverpool(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Liverpool")
	defer span.End()
}

func edinburgh(ctx context.Context) {
	_, span := otel.Tracer.Start(ctx, "Edinburgh")
	defer span.End()
}
