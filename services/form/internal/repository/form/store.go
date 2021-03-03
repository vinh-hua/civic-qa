package form

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	ErrFormNotFound = errors.New("Form Does Not Exist")
	ErrUserNotFound = errors.New("User Does Not Exist")
)

type Store interface {
	Create(form *model.Form) error
	GetByID(formID uint) (*model.Form, error)
	GetByUserID(userID uint) ([]*model.Form, error)
}
