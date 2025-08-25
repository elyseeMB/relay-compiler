package main

import (
	"context"
	"log"
	"net/http"
	"os"

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

	serverHanlder, err := server.NewServer()

	if err != nil {
		panic(err)
	}

	impl := db.New()
	unit := unit.NewUnit(impl, "tp", "v1", "dev")

	err = unit.Run()

	if err != nil && err != context.Canceled {
		panic(err)
	}

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":8080", serverHanlder))

}
