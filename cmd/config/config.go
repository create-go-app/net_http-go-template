package config

import (
	"os"
	"strconv"
	"strings"
)

// Config struct for app config
type Config struct {
	Server            server
	DebugMode         bool
	FrontendBuildPath string
}

type server struct {
	Host         string
	Port         string
	WriteTimeout int
	ReadTimeout  int
	IdleTimeout  int
}

// Load returns a new Config struct
func Load() *Config {
	return &Config{
		Server: server{
			Host:         getEnv("HOST", "127.0.0.1"),
			Port:         getEnv("PORT", "8080"),
			ReadTimeout:  getEnvAsInt("READ_TIMEOUT", 5),
			WriteTimeout: getEnvAsInt("WRITE_TIMEOUT", 10),
			IdleTimeout:  getEnvAsInt("IDLE_TIMEOUT", 15),
		},
		DebugMode:         getEnvAsBool("DEBUG_MODE", true),
		FrontendBuildPath: getEnv("FRONTEND_BUILD_PATH", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultValue int) int {
	checkValue := getEnv(name, "")

	if value, err := strconv.Atoi(checkValue); err == nil {
		return value
	}

	return defaultValue
}

// Helper to read an environment variable into a bool or return default value
func getEnvAsBool(name string, defaultValue bool) bool {
	checkValue := getEnv(name, "")

	if val, err := strconv.ParseBool(checkValue); err == nil {
		return val
	}

	return defaultValue
}

// Helper to read an environment variable into a string slice or return default value
func getEnvAsSlice(name string, defaultValue []string, sep string) []string {
	checkValue := getEnv(name, "")

	if checkValue == "" {
		return defaultValue
	}

	val := strings.Split(checkValue, sep)

	return val
}
