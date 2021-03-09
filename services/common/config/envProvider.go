package config

import (
	"fmt"
	"os"
)

// EnvProvider provides configuration from environment variables
type EnvProvider struct {
	_         Provider
	isVerbose bool
}

// Name returns the provider name
func (EnvProvider) Name() string {
	return "environment"
}

// SetVerbose sets whether this Provider should log results
func (e *EnvProvider) SetVerbose(set bool) {
	e.isVerbose = set
}

// verbose returns whether this Provider should log results
func (e *EnvProvider) verbose() bool {
	return e.isVerbose
}

// GetOrFallback returns a key if found, or a fallback
func (e *EnvProvider) GetOrFallback(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		verboseFallback(e, key, fallback)
		return fallback
	}
	verboseGet(e, key, value)
	return value
}

// Get returns a key or an error
func (e *EnvProvider) Get(key string) (string, KeyError) {
	value := os.Getenv(key)
	if len(value) == 0 {
		return "", configError(fmt.Sprintf("Key not found: %s", key))
	}
	verboseGet(e, key, value)
	return value, nil
}
