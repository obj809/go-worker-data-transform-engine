package main

import (
	"log"
	"net/http"
	"os"

	"github.com/softdev/go-worker-data-transform-engine/handlers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlers.HealthHandler)
	mux.HandleFunc("/process", handlers.ProcessHandler)

	addr := ":" + port
	log.Printf("Go worker starting on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
