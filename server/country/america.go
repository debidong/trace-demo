package country

import (
	"net/http"
	"trace-demo/server/otel"
)

func America(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "America")
	defer span.End()

	americaPool.GetRandomCity()(w, r)
}

// new-york -> los-angeles -> chicago
func NewYork(w http.ResponseWriter, r *http.Request) {
	newyork(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New York"))
}

// los-angeles -> chicago
func LosAngeles(w http.ResponseWriter, r *http.Request) {
	losangeles(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Los Angeles"))
}

func Chicago(w http.ResponseWriter, r *http.Request) {
	chicago(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Chicago"))
}

// to england
func Edmond(w http.ResponseWriter, r *http.Request) {
	edmond(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Edmond"))
}

func newyork(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "New York")
	defer span.End()

	losangeles(w, r)
}

func losangeles(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Los Angeles")
	defer span.End()

	chicago(w, r)
}

func chicago(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Chicago")
	defer span.End()
}

func edmond(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Edmond")
	defer span.End()

	England(w, r)
}
