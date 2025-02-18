package main

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"sync"
	"trace-demo/requester"
	"trace-demo/server"
	"trace-demo/server/otel"
)

var countries = []string{"china", "america", "england"}

func main() {
	// Handle SIGINT (CTRL+C) gracefully.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// Set up OpenTelemetry.
	otelShutdown, err := otel.SetupOTelSDK(ctx)
	if err != nil {
		return
	}
	// Handle shutdown properly so nothing leaks.
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	srv := server.NewServer(ctx)
	srvErr := make(chan error, 1)
	go func() {
		srvErr <- srv.ListenAndServe()
	}()

	requesters := make(map[string]*requester.Requester)
	var wg sync.WaitGroup
	wg.Add(len(countries))
	for _, country := range countries {
		go func(country string) {
			defer wg.Done()
			requester := requester.NewRequester(country)
			requesters[country] = requester
			requester.StartRequest()
			log.Default().Println("main: started requester for", country)
		}(country)
	}
	wg.Wait()

	select {
	case err := <-srvErr:
		log.Default().Println("main: server error", err)
	case <-ctx.Done():
		log.Default().Println("main: context done")
	}

	for country, requester := range requesters {
		requester.StopRequest()
		log.Default().Println("main: stopped requester for", country)
	}
	log.Default().Println("main: program terminated")
}
