package apiserver

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// APIServer struct
type APIServer struct {
	config *Config
	logger *zap.Logger
	router *mux.Router
}

// New method for init new server instance
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: Logger(config),
		router: mux.NewRouter(),
	}
}

// Start method for start new server
func (s *APIServer) Start() error {
	// Starting message
	s.logger.Info(
		"Starting API server",
		zap.String("host", s.config.Server.Host),
		zap.String("port", s.config.Server.Port),
	)

	// Init router
	s.Router()

	// Define server options
	server := &http.Server{
		Addr:         s.config.Server.Host + ":" + s.config.Server.Port,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.config.Server.Timeout.Read) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.Timeout.Write) * time.Second,
		IdleTimeout:  time.Duration(s.config.Server.Timeout.Idle) * time.Second,
	}

	// Start server
	return server.ListenAndServe()
}
