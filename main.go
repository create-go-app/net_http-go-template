package main

import (
	"github.com/create-go-app/net_http-go-template/pkg/apiserver"
)

func main() {
	// Parse config path from environment variable.
	configPath := apiserver.GetEnv("CONFIG_PATH", "configs/apiserver.yml")

	// Create new config.
	config, err := apiserver.NewConfig(configPath)
	apiserver.ErrChecker(err)

	// Create new server.
	server := apiserver.NewServer(config)

	// Start API server.
	server.Start()
}
