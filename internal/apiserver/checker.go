package apiserver

import (
	"go.uber.org/zap"
	"os"
)

// IsError function for check error and show message
func (s *APIServer) IsError(err error) {
	// If got error
	if err != nil {
		// Show zap logger error
		s.logger.Error("Error", zap.Error(err))

		// Exit with status 1
		os.Exit(1)
	}
}
