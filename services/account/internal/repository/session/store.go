// Package session defines the session.Store interface
// as well as its implementations.
package session

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/team-ravl/civic-qa/services/common/model"
)

// Token is a type alias for string
// represents a users session
type Token string

const (
	// Number of bytes in generated tokens
	tokenSize = 64
	// InvalidSessionToken is returned alongside errors
	InvalidSessionToken Token = "INVALID"
)

var (
	// ErrStateNotFound is returned when a requested state doesn't exist
	ErrStateNotFound = errors.New("State Does Not Exist")
)

// Store interface describes implementations of user session storage
type Store interface {
	// Create starts a session
	Create(state model.SessionState) (Token, error)
	// Get retrieves a session
	Get(token Token) (*model.SessionState, error)
	// Delete ends a Session
	Delete(token Token) error
}

// generateToken returns a base64 URL encoded
// SessionToken of length const tokenSize
func generateToken() (Token, error) {
	token := make([]byte, tokenSize)
	_, err := rand.Read(token)
	if err != nil {
		return InvalidSessionToken, err
	}
	return Token(base64.URLEncoding.EncodeToString(token)), nil
}
