package country

import (
	"net/http"

	"go.opentelemetry.io/otel"
)

func England(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("england")
	_, span := tracer.Start(r.Context(), "England")
	defer span.End()

	englandPool.GetRandomCity()(w, r)
}

// london -> manchester -> liverpool
func London(w http.ResponseWriter, r *http.Request) {
	london(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("London"))
}

// manchester -> liverpool
func Manchester(w http.ResponseWriter, r *http.Request) {
	manchester(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Manchester"))
}

func Liverpool(w http.ResponseWriter, r *http.Request) {
	liverpool(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Liverpool"))
}

// to china
func Edinburgh(w http.ResponseWriter, r *http.Request) {
	edinburgh(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edinburgh"))
}

func london(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("england")
	_, span := tracer.Start(r.Context(), "London")
	defer span.End()

	manchester(w, r)
}

func manchester(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("england")
	_, span := tracer.Start(r.Context(), "Manchester")
	defer span.End()

	liverpool(w, r)
}

func liverpool(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("england")
	_, span := tracer.Start(r.Context(), "Liverpool")
	defer span.End()
}

func edinburgh(w http.ResponseWriter, r *http.Request) {
	tracer := otel.Tracer("england")
	_, span := tracer.Start(r.Context(), "Edinburgh")
	defer span.End()

	China(w, r)
}
