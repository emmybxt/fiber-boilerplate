package action

import (
	"encoding/json"
	"net/http"
)

// HealthCheck is a simple HTTP handler that responds with a JSON
// object indicating the server is healthy.
func HealthCheck(w http.ResponseWriter, _ *http.Request) {
	response := map[string]interface{}{
		"status": "ok",
		"code":   http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
