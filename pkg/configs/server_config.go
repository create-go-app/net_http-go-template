package configs

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

// ServerConfig func for configuration net/http app.
func ServerConfig(router *mux.Router) *http.Server {
	return &http.Server{
		Handler:      router,
		Addr:         os.Getenv("SERVER_URL"),
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}
