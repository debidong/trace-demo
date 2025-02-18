package country

import (
	"net/http"
	"trace-demo/server/otel"
)

func China(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "China")
	defer span.End()

	chinaPool.GetRandomCity()(w, r)
}

// beijing -> shanghai -> guangzhou
func Beijing(w http.ResponseWriter, r *http.Request) {
	beijing(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Beijing"))
}

// shanghai -> guangzhou
func Shanghai(w http.ResponseWriter, r *http.Request) {
	shanghai(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Shanghai"))
}

func Guangzhou(w http.ResponseWriter, r *http.Request) {
	guangzhou(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Guangzhou"))
}

// to america
func Handan(w http.ResponseWriter, r *http.Request) {
	handan(w, r)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Handan"))
}

func beijing(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Beijing")
	defer span.End()

	shanghai(w, r)
}

func shanghai(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Shanghai")
	defer span.End()

	guangzhou(w, r)
}

func guangzhou(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Guangzhou")
	defer span.End()
}

func handan(w http.ResponseWriter, r *http.Request) {
	_, span := otel.Tracer.Start(r.Context(), "Handan")
	defer span.End()

	America(w, r)
}
