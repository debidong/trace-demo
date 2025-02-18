package country

import (
	"net/http"
	"trace-demo/server/otel"
)

func England(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "England Custom")
	r.WithContext(ctx)
	defer span.End()

	englandPool.GetRandomCity()(w, r)
}

// london -> manchester -> liverpool
func London(w http.ResponseWriter, r *http.Request) {
	london(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("London"))
}

// manchester -> liverpool
func Manchester(w http.ResponseWriter, r *http.Request) {
	manchester(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Manchester"))
}

func Liverpool(w http.ResponseWriter, r *http.Request) {
	liverpool(r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Liverpool"))
}

// to china
func Edinburgh(w http.ResponseWriter, r *http.Request) {
	edinburgh(r)
	China(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edinburgh"))
}

func london(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "London")
	r.WithContext(ctx)
	defer span.End()

	manchester(r)
}

func manchester(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Manchester")
	r.WithContext(ctx)
	defer span.End()

	liverpool(r)
}

func liverpool(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Liverpool")
	r.WithContext(ctx)
	defer span.End()
}

func edinburgh(r *http.Request) {
	ctx, span := otel.Tracer.Start(r.Context(), "Edinburgh")
	r.WithContext(ctx)
	defer span.End()
}
