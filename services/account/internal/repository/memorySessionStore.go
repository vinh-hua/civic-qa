package repository

import (
	"crypto/rand"
	"encoding/base64"
	"sync"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
)

const tokenSize = 64

type MemorySessionStore struct {
	data map[SessionToken]model.SessionState
	lock sync.Mutex
}

func NewMemorySessionStore(intialCapacity int) *MemorySessionStore {
	return &MemorySessionStore{data: make(map[SessionToken]model.SessionState, intialCapacity)}
}

func (s *MemorySessionStore) Create(state model.SessionState) (SessionToken, error) {
	s.lock.Lock()
	token, err := generateToken()
	if err != nil {
		return InvalidSessionToken, err
	}
	s.data[token] = state
	s.lock.Unlock()
	return token, nil
}

func (s *MemorySessionStore) Get(token SessionToken) (*model.SessionState, error) {
	s.lock.Lock()
	state, ok := s.data[token]
	if !ok {
		return nil, ErrStateNotFound
	}

	s.lock.Unlock()
	return &state, nil
}

func (s *MemorySessionStore) Delete(token SessionToken) error {
	s.lock.Lock()
	delete(s.data, token)
	s.lock.Unlock()
	return nil
}

func generateToken() (SessionToken, error) {
	token := make([]byte, tokenSize)
	_, err := rand.Read(token)
	if err != nil {
		return InvalidSessionToken, err
	}
	return SessionToken(base64.URLEncoding.EncodeToString(token)), nil
}
