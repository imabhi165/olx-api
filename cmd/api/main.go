package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/imabhi165/olx-api/internal/config"
	"github.com/imabhi165/olx-api/internal/db"
	"github.com/imabhi165/olx-api/internal/handlers"
)

// main configures and starts the HTTP server.
//
// It registers the /healthz route on a fresh ServeMux, constructs an
// http.Server with explicit read, write, and idle timeouts, and then
// blocks in ListenAndServe on :8090. If the server fails to start or
// stops unexpectedly, the process exits with a fatal log entry.
func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DatabaseUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}
	fmt.Println("Database Connected")
	fmt.Println("Starting Olx server...")

	//
	lh := handlers.NewListingHandlers(db)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", handlers.HealthzHandler)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	log.Printf("Server is listening on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
