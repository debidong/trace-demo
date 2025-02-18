package server

import (
	"context"
	"net"
	"net/http"
	"time"
)

func NewServer(ctx context.Context) *http.Server {
	// Start HTTP server.
	return &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      NewHTTPHandler(),
	}
}
