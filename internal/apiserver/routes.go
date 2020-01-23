package apiserver

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gorilla/mux"
	jsoniter "github.com/json-iterator/go"
)

// Define JSON iterator
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// frontendHandler implements the http.Handler interface, so we can use it
// to respond to HTTP requests. The path to the static directory and
// path to the index file within that static directory are used to
// serve the SPA in the given static directory.
type frontendHandler struct {
	staticPath string
	indexPath  string
}

// Router method for create router
func (s *APIServer) Router() {
	// API routes with allowed methods
	s.router.HandleFunc("/api/index", s.handleIndex()).Methods(http.MethodGet)

	// Frontend route
	if s.config.Static.Path != "" {
		frontend := &frontendHandler{
			staticPath: s.config.Static.Path,
			indexPath:  "index.html",
		}
		s.router.PathPrefix("/").Handler(frontend)
	}

	// Middlewares
	s.router.Use(
		mux.CORSMethodMiddleware(s.router), // CORS
		s.loggerMiddleware,                 // logger
	)
}

// handleIndex method for handle /api/index route
func (s *APIServer) handleIndex() http.HandlerFunc {
	// Optional struct
	type request struct{}

	// Return func
	return func(w http.ResponseWriter, r *http.Request) {
		// Set headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Content-Type", "application/json")

		// Return JSON
		json.NewEncoder(w).Encode(
			map[string]bool{
				"ok": true,
			},
		)
	}
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (f *frontendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// If we failed to get the absolute path respond
		// with a 400 bad request and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Prepend the path with the path to the static directory
	path = filepath.Join(f.staticPath, path)

	// Check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// File does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(f.staticPath, f.indexPath))
		return
	} else if err != nil {
		// If we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(f.staticPath)).ServeHTTP(w, r)
}
