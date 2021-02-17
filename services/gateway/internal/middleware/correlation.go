package middleware

import (
	"log"
	"net/http"

	"github.com/google/uuid"
)

const (
	correlationHeader = "X-Correlation-ID"
	noCorrelationID   = "NO-CORRELATION-ID"
)

// generateCorrelationID generates a new UUIDv4 and returns it as a string
func generateCorrelationID() (string, error) {
	uuid4, err := uuid.NewRandom()
	if err != nil {
		return noCorrelationID, err
	}

	return uuid4.String(), nil
}

// NewCorrelatorMiddleware is a middleware that adds a randomly generated uuid
// to each request header, intended to help cross-service tracking
func NewCorrelatorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID, err := generateCorrelationID()
		if err != nil {
			log.Printf("Error generating correlationID: %v", err)
		} else {
			r.Header.Set(correlationHeader, correlationID)
		}
		next.ServeHTTP(w, r)
	})
}
