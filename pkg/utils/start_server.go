package utils

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// StartServerWithGracefulShutdown function for starting server with a graceful shutdown.
func StartServerWithGracefulShutdown(server *http.Server) {
	// Define waiting time.
	var wait time.Duration

	// Create channel for idle connections.
	idleConnsClosed := make(chan struct{})

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt) // catch OS signals
		<-sigint

		// Received an interrupt signal, shutdown.
		if err := server.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("Oops... Server is not shutting down! Reason: %v", err)
		}

		close(idleConnsClosed)
	}()

	// Run server.
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Oops... Server is not running! Reason: %v", err)
	}

	<-idleConnsClosed
}
