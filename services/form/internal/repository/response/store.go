package response

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	// ErrResponseNotFound is returned when a requested FormResponse does not exist
	ErrResponseNotFound = errors.New("Response Does Not Exist")
)

// Store is an interface for implementations of response.Store.
// Describes storage of FormResponses
type Store interface {
	Create(response *model.FormResponse) error
	GetByID(responseID uint) (*model.FormResponse, error)
	GetByFormID(formID uint) ([]*model.FormResponse, error)
	GetByUserID(userID uint) ([]*model.FormResponse, error)
}
