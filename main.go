package main

import (
	"github.com/create-go-app/net_http-go-template/internal/apiserver"
	"github.com/create-go-app/net_http-go-template/internal/checker"
	"github.com/joho/godotenv"
)

func init() {
	// Load values from .env
	err := godotenv.Load()
	checker.IsError(err, "Not found .env file in project root folder!")
}

func main() {
	// Create new config
	config := apiserver.NewConfig()

	// Create new server
	server := apiserver.New(config)
	checker.IsError(server.Start())
}
