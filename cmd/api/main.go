package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// healthzHandler handles GET /healthz requests.
// It sets the response content type to JSON, writes a 200 OK status,
// and returns a small JSON payload indicating the service is healthy.
// Method and path matching are enforced by the ServeMux pattern
// "GET /healthz", so this handler is only invoked for that route.
func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status": "OK"}`))
}

// main configures and starts the HTTP server.
//
// It registers the /healthz route on a fresh ServeMux, constructs an
// http.Server with explicit read, write, and idle timeouts, and then
// blocks in ListenAndServe on :8090. If the server fails to start or
// stops unexpectedly, the process exits with a fatal log entry.
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthzHandler)

	srv := http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
