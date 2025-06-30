package handlers

import (
	"net/http"
)

// Add your handler functions here.

// ExampleHandler is a placeholder HTTP handler.
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from auth handler!"))
}
