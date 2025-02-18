package server

import (
	"context"
	"net"
	"net/http"
	"time"
	"trace-demo/config"
)

func NewServer(ctx context.Context) *http.Server {
	// Start HTTP server.
	return &http.Server{
		Addr:         config.ServerAddr,
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      NewHTTPHandler(),
	}
}
