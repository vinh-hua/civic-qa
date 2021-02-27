// Package config defines the Provider interface and implementations.
// config.Provider implementors are used to load application configuration
// from various sources with a standardized interface.
package config

import (
	"errors"
	"log"
)

// KeyError is returned when Get cannot provide a key
type KeyError error

func configError(msg string) KeyError {
	return errors.New(msg)
}

// Provider describes implementations of configuration providers
type Provider interface {
	Name() string
	SetVerbose(set bool)
	GetOrFallback(key, fallback string) string
	Get(key string) (string, KeyError)
	verbose() bool
}

func verboseGet(provider Provider, key, value string) {
	if provider.verbose() {
		log.Printf("Provider %s: key %s set to %s", provider.Name(), key, value)
	}
}

func verboseFallback(provider Provider, key, fallback string) {
	if provider.verbose() {
		log.Printf("Provider %s: key %s not set, fallback: %s", provider.Name(), key, fallback)
	}
}
