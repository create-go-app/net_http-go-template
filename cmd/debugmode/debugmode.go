package debugmode

import (
	"net/http"

	"github.com/create-go-app/net_http-go-template/cmd/checker"
	"go.uber.org/zap"
)

// Enable function for show debug info
func Enable(next http.Handler, debug bool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If debug is true, show messages
		if debug == true {
			// Init zap logger
			logger, err := zap.NewDevelopment()
			checker.IsError(err)
			defer logger.Sync()

			// Show report for each request
			logger.Debug("fetching URL",
				zap.String("method", r.Method),
				zap.String("url", r.RequestURI),
			)
		}

		// Call the next handler, which can be another middleware
		// in the chain, or the final handler
		next.ServeHTTP(w, r)
	})
}
