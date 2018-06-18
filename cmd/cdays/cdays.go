package main

import (
	"log"
	"os"

	"net/http"

	"os/signal"
	"syscall"

	"context"
	"time"

	"github.com/jumale/cdays/internal/routing"
	"github.com/jumale/cdays/internal/version"
)

func main() {
	log.Printf(
		"The app is starting, version is %s, build time is %s, commit is %s...",
		version.Release,
		version.BuildTime,
		version.Commit,
	)

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}
	diagPort := os.Getenv("DIAGNOSTICS_PORT")
	if port == "" {
		log.Fatal("The port wasn't set")
	}

	var blServer, diagServer http.Server

	srvErrs := make(chan error, 2)

	go func() {
		blServer = http.Server{
			Addr:    ":" + port,
			Handler: routing.NewBLRouter(),
		}
		err := blServer.ListenAndServe()
		srvErrs <- err
	}()

	go func() {
		diagServer = http.Server{
			Addr:    ":" + diagPort,
			Handler: routing.NewDiagnosticsRouter(),
		}
		err := diagServer.ListenAndServe()
		srvErrs <- err
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	select {
	case killSignal := <-interrupt:
		log.Printf("Got %s. Stopping...", killSignal)
	case err := <-srvErrs:
		log.Printf("Got a server error: %s", err)
	}

	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		blServer.Shutdown(ctx)
	}

	{
		ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelFunc()
		diagServer.Shutdown(ctx)
	}
}
