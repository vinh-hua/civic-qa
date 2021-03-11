// Package user defines the user.Store interface
// as well as its implementations.
package user

import (
	"errors"

	"github.com/team-ravl/civic-qa/services/common/model"
)

var (
	// ErrUserNotFound is returned when an operation is attempted on a user
	// that does not exist
	ErrUserNotFound = errors.New("User Does Not Exist")
)

// Store interface describes implementations of user storage
type Store interface {
	Create(user *model.User) error
	GetByID(id uint) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	EmailInUse(email string) (bool, error)
}
