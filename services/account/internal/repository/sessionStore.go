package repository

import (
	"errors"

	"github.com/vivian-hua/civic-qa/service/account/internal/model"
)

type SessionToken string

const InvalidSessionToken SessionToken = "INVALID"

var (
	ErrStateNotFound = errors.New("State Does Not Exist")
)

type SessionStore interface {
	Create(state model.SessionState) (SessionToken, error)
	Get(token SessionToken) (*model.SessionState, error)
	Delete(token SessionToken) error
}
