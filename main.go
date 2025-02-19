package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"os"
	"os/signal"
	"trace-demo/logic"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	srvName := flag.String("n", "", "name of the server")
	cfgPath := flag.String("c", "", "path to the config file")
	flag.Parse()

	if *cfgPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	if *srvName == "" {
		flag.Usage()
		os.Exit(1)
	}

	logic.MustLoadConfig(*cfgPath)
	otelShutdown, err := logic.SetupOTelSDK(ctx, *srvName)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	srv := logic.NewServer(*srvName)
	errSrv := make(chan error)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			errSrv <- err
		}
	}()

	requester := logic.NewRequester(*srvName)
	requester.StartRequest()

	defer func() {
		if err != nil {
			log.Default().Fatal(err)
		}
		requester.StopRequest()
		srv.Shutdown(context.Background())
	}()

	select {
	case _err := <-errSrv:
		errors.Join(err, _err)
	case <-ctx.Done():
		log.Default().Println("main: context done")
	}
}
