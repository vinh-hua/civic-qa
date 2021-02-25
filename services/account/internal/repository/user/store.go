package user

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/common/model"
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
}
