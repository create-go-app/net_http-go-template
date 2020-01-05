package apiserver

import (
	"net/http"
	"os"
	"time"

	"github.com/create-go-app/net_http-go-template/internal/checker"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// APIServer ...
type APIServer struct {
	config *Config
	logger *zap.Logger
	router *mux.Router
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: zap.New(zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.Lock(os.Stdout),
			zap.NewAtomicLevel(),
		)),
		router: mux.NewRouter(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	// Init config for logger
	checker.IsError(s.configureLogger())

	// Init config for router
	s.configureRouter()

	// Starting message
	s.logger.Info(
		"Starting API server",
		zap.String("host", s.config.Server.Host),
		zap.String("port", s.config.Server.Port),
	)

	// If app have frontend folder with builded bundle
	// if _, err := os.Stat(c.FrontendBuildPath); !os.IsNotExist(err) {
	// 	r.PathPrefix("/").Handler(http.StripPrefix(
	// 		c.FrontendBuildPath,
	// 		http.FileServer(http.Dir(c.FrontendBuildPath)),
	// 	))
	// }

	// Define server options
	server := &http.Server{
		Addr:         s.config.Server.Host + ":" + s.config.Server.Port,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.config.Server.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.config.Server.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(s.config.Server.IdleTimeout) * time.Second,
	}

	// Start server
	return server.ListenAndServe()
}

// configureLogger ...
func (s *APIServer) configureLogger() error {
	// Define log level
	level := zap.NewAtomicLevel()

	// Set log level from .env file
	switch s.config.LogLevel {
	case "debug":
		level.SetLevel(zap.DebugLevel)
	case "warn":
		level.SetLevel(zap.WarnLevel)
	case "error":
		level.SetLevel(zap.ErrorLevel)
	case "fatal":
		level.SetLevel(zap.FatalLevel)
	case "panic":
		level.SetLevel(zap.PanicLevel)
	default:
		level.SetLevel(zap.InfoLevel)
	}

	// Config for zap output
	// encoderCfg := zap.NewProductionEncoderConfig()

	// // Formated timestamp in the output.
	// encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
	// zap.NewProductionEncoderConfig().EncodeTime = zapcore.RFC3339TimeEncoder

	// // Create new zap logger
	// logger := zap.New(zapcore.NewCore(
	// 	zapcore.NewJSONEncoder(encoderCfg),
	// 	zapcore.Lock(os.Stdout),
	// 	level,
	// ))
	// defer logger.Sync()

	return nil
}

// configureRouter ...
func (s *APIServer) configureRouter() {
	// API Index route
	s.router.HandleFunc("/api/index", s.handleIndex())
}
