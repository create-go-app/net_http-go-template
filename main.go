package main

import (
	"github.com/create-go-app/net_http-go-template/pkg/configs"
	"github.com/create-go-app/net_http-go-template/pkg/utils"
	"github.com/gorilla/mux"
)

func main() {
	// Initialize a new router.
	router := mux.NewRouter()

	// Register API routes.
	server := configs.ServerConfig(router)

	// Start API server.
	utils.StartServerWithGracefulShutdown(server)
}
