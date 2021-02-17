package middleware

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

// NewLoggingMiddleware returns an http.Handler middleware to log all
// requests before forwarding them to the next handler
func NewLoggingMiddleware(output *os.File, aggregatorAddress string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return handlers.LoggingHandler(output, next)
	}
}
