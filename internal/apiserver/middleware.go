package apiserver

import (
	"net/http"

	"go.uber.org/zap"
)

// loggerMiddleware function middleware for logging
func (s *APIServer) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Show debug message
		s.logger.Debug(
			"fetch URL",
			zap.String("method", r.Method),
			zap.String("url", r.URL.Path),
		)

		// Call the next handler, which can be another middleware
		// in the chain, or the final handler
		next.ServeHTTP(w, r)
	})
}
