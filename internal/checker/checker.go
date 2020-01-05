package checker

import (
	"log"
	"os"
)

const (
	// Color:
	red string = "\033[0;31m"

	// Clear color
	noColor string = "\033[0m"
)

// IsError function for check error and show message
func IsError(err error, msg ...string) {
	if err != nil {
		// Show error report
		if msg != nil {
			// Custom error message
			log.Printf("%v[✘] Error: %v%v\n", red, msg, noColor)
		} else {
			// Else no custom message, show default error message
			log.Printf("%v[✘] Error: %v%v\n", red, err, noColor)
		}

		// Exit with status 1
		os.Exit(1)
	}
}
