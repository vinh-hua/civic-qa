package environment

import (
	"log"
	"os"
)

// GetEnvOrFallback returns the value of environment
// variable key, or returns fallback if key is unset/empty
// after logging the decisions
func GetEnvOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Printf("Environment variable \"%s\" was empty, fallback: \"%s\"", key, fallback)
		return fallback
	}

	log.Printf("Environment variable \"%s\" was set to: \"%s\"", key, value)
	return value
}

// GetEnvOrFatal returns the value of environment
// variable key, or raises fatal,
// logs the decsision
func GetEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatalf("Environment variable \"%s\" was empty, fatal.", key)
	}

	log.Printf("Environment variable \"%s\" was set to: \"%s\".", key, value)
	return value
}
