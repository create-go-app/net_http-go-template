package apiserver

import (
	"net/http"
)

// handleIndex ...
func (s *APIServer) handleIndex() http.HandlerFunc {
	// Optional struct
	type request struct{}

	// Return func
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.URL.Path))
	}
}
