package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elyseeMB/relay-compiler/pkg/db"
	"github.com/elyseeMB/relay-compiler/pkg/server"
	"go.gearno.de/kit/unit"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	serverHandler, err := server.NewServer()
	if err != nil {
		panic(err)
	}

	impl := db.New()
	unitInstance := unit.NewUnit(impl, "tp", "v1", "dev")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := unitInstance.Run()
		if err != nil && err != context.Canceled {
			log.Printf("Unit error: %v", err)
			cancel()
		}
	}()

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: serverHandler,
	}

	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server error: %v", err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-quit:
		log.Println("Shutting down...")
	case <-ctx.Done():
		log.Println("Context canceled...")
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	srv.Shutdown(shutdownCtx)
	cancel()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
}
