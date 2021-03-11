package tag

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	// ErrTagNotFound is returned when a requested tag does not exist
	ErrTagNotFound = errors.New("Tag Does Not Exist")
)

// Store is an interface for implementations of tag.Store.
// Describes storage of tags
type Store interface {
	GetAll(userID uint) ([]*model.Tag, error)
	GetByResponseID(userID uint, responseID uint) ([]*model.Tag, error)
	Create(userID uint, responseID uint, value string) error
	Delete(userID uint, responseID uint, value string) error
}
