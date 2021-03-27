package configs

import (
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/create-go-app/net_http-go-template/pkg/utils"
	"github.com/gorilla/mux"
)

// ServerConfig func for configuration net/http app.
func ServerConfig(router *mux.Router) *http.Server {
	// Define server settings:
	serverConnURL, _ := utils.ConnectionURLBuilder("server")
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// Return server configuration.
	return &http.Server{
		Handler:     router,
		Addr:        serverConnURL,
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
