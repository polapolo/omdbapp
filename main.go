package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	httpserver "github.com/polapolo/omdbapp/app/http"
)

func main() {
	var wait time.Duration

	// HTTP Server
	httpServer := httpserver.InitHTTPServer()
	go func() {
		log.Println("[omdbapp][API] Served on " + httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	httpServer.Shutdown(ctx)
	// Optionally, you could run httpServer.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("shutting down")
	os.Exit(0)
}
