package handlers

import "net/http"

// StatsHandler is a sample HTTP handler function.
func StatsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Stats OK"))
}
