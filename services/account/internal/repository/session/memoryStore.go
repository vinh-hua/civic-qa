package session

import (
	"sync"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
)

const intialCapacity = 100

// MemoryStore implements SessionStore,
// Threadsafe via RWMutex
type MemoryStore struct {
	data map[Token]model.SessionState
	lock sync.RWMutex
}

// NewMemoryStore returns a MemorySessionStore
func NewMemoryStore() *MemoryStore {
	return &MemoryStore{data: make(map[Token]model.SessionState, intialCapacity)}
}

// Create begins a new session for a given SessionState, returning a new SessionToken
func (s *MemoryStore) Create(state model.SessionState) (Token, error) {
	s.lock.Lock()
	// generate a new token as the key for the state
	token, err := generateToken()
	if err != nil {
		return InvalidSessionToken, err
	}
	s.data[token] = state
	s.lock.Unlock()
	return token, nil
}

// Get returns a SessionState for a given SessionToken if it exists
func (s *MemoryStore) Get(token Token) (*model.SessionState, error) {
	s.lock.RLock()
	state, ok := s.data[token]
	if !ok {
		return nil, ErrStateNotFound
	}

	s.lock.RUnlock()
	return &state, nil
}

// Delete removes a SessionState given a SessionToken.
// error is always nil in this implementation
func (s *MemoryStore) Delete(token Token) error {
	s.lock.Lock()
	delete(s.data, token)
	s.lock.Unlock()
	return nil
}
