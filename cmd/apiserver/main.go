package main

import (
	"github.com/create-go-app/net_http-go-template/internal/apiserver"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// Load values from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create new config
	config := apiserver.NewConfig()

	// Create new server
	server := apiserver.New(config)

	// Start server
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
