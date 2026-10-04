package handlers

import "net/http"

// healthzHandler handles GET /healthz requests.
// It sets the response content type to JSON, writes a 200 OK status,
// and returns a small JSON payload indicating the service is healthy.
// Method and path matching are enforced by the ServeMux pattern
// "GET /healthz", so this handler is only invoked for that route.
func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(`{"status": "OK"}`))
}
