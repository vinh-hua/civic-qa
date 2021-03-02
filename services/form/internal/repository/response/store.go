package response

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
	GetByID(id uint) (*model.FormResponse, error)
	GetByFormID(formID uint) ([]*model.FormResponse, error)
}
