package form

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	// ErrFormNotFound is returned when a requested form does not exist
	ErrFormNotFound = errors.New("Form Does Not Exist")
)

// Store is an interface for implementations of form.Store.
// Describes storage of forms
type Store interface {
	Create(form *model.Form) error
	GetByID(formID uint) (*model.Form, error)
	GetByUserID(userID uint) ([]*model.Form, error)
}
