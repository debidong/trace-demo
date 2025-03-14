package logic

import (
	"math/rand"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

const (
	gormDBPath = "gorm.db"
)

var urls []string

func init() {
	urls = make([]string, 0)
	for srvName, srv := range MustLoadConfig("config.yaml").Server {
		for _, uri := range srv.Uri {
			urls = append(urls, FormatRequestURL(srvName, uri))
		}
	}
}

func handleCommon(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), r.URL.Path)
	defer span.End()

	// whether to return immediately
	if rand.Intn(2) == 0 {
		return
	}

	// get resource
	next := rand.Intn(len(urls))
	req, _ := http.NewRequestWithContext(ctx, "GET", urls[next], nil)

	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, propagation.HeaderCarrier(req.Header))

	client := &http.Client{
		Transport: http.DefaultTransport,
	}
	client.Do(req)
	return
}

func handleGorm(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(sqlite.Open(gormDBPath), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	ctx := r.Context()
	// db.Use(tracing.NewPlugin(tracing.WithoutMetrics()))
	db.Use(tracing.NewPlugin(tracing.WithTracerProvider(otel.GetTracerProvider())))
	query := "SELECT " + "\"" + r.URL.Path + "\""
	db.WithContext(ctx).Exec(query)
	return
}
