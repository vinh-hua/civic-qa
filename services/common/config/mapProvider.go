package config

import "fmt"

// MapProvider implements Provider
type MapProvider struct {
	_         Provider
	Data      map[string]string
	isVerbose bool
}

// Name returns the Provider name
func (MapProvider) Name() string {
	return "map"
}

// SetVerbose sets whether this Provider should log results
func (m *MapProvider) SetVerbose(set bool) {
	m.isVerbose = set
}

// verbose returns whether this Provider should log results
func (m *MapProvider) verbose() bool {
	return m.isVerbose
}

// GetOrFallback returns a key or a fallback
func (m *MapProvider) GetOrFallback(key, fallback string) string {
	value, ok := m.Data[key]
	if !ok {
		verboseFallback(m, key, fallback)
		return fallback
	}
	verboseGet(m, key, value)
	return value
}

// Get returns a key or an Error
func (m *MapProvider) Get(key string) (string, KeyError) {
	value, ok := m.Data[key]
	if !ok {
		return "", configError(fmt.Sprintf("Key not found: %s", key))
	}
	verboseGet(m, key, value)
	return value, nil
}
