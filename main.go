package main

import (
	"context"
	"log"
	"lumoshive-be-chap38-39/infra"
	"lumoshive-be-chap38-39/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Initialize the database connection
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal(err)
	}

	r := router.NewRouter(*ctx)

	// Start the server on the configured port
	if ctx.Cfg.AppPort == "" {
		log.Fatal("PORT is not defined in the configuration")
	}

	srv := &http.Server{
		Addr:    ":" + ctx.Cfg.AppPort,
		Handler: r,
	}

	go func() {
		// Start the server
		log.Printf("Server running on port %s", ctx.Cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	// Create a timeout context for graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shutdown the server
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// Catching context timeout
	select {
	case <-shutdownCtx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server exiting")
}
