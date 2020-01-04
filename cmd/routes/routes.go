package routes

import (
	"net/http"
)

// Index ...
func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.Path))
}
