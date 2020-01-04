package main

import (
	"net/http"
	"os"
	"time"

	"github.com/create-go-app/net_http-go-template/cmd/checker"
	"github.com/create-go-app/net_http-go-template/cmd/config"
	"github.com/create-go-app/net_http-go-template/cmd/debugmode"
	"github.com/create-go-app/net_http-go-template/cmd/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// Load values from .env
	err := godotenv.Load()
	checker.IsError(err, "Not found .env file in project root folder!")
}

func main() {
	// Load config
	c := config.Load()

	// Define gorilla/mux router and app routes
	r := mux.NewRouter()
	r.HandleFunc("/api/index", routes.Index)

	// If app have frontend folder with builded bundle
	if _, err := os.Stat(c.FrontendBuildPath); !os.IsNotExist(err) {
		r.PathPrefix("/").Handler(http.StripPrefix(
			c.FrontendBuildPath,
			http.FileServer(http.Dir(c.FrontendBuildPath)),
		))
	}

	// Define net/http server with options
	server := &http.Server{
		Addr:         c.Server.Host + ":" + c.Server.Port,
		Handler:      debugmode.Enable(r, c.DebugMode),
		ReadTimeout:  time.Duration(c.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(c.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(c.Server.IdleTimeout) * time.Second,
	}

	// Start server
	checker.IsError(server.ListenAndServe())
}
