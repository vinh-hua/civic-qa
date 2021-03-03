package response

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

var (
	ErrResponseNotFound = errors.New("Response Does Not Exist")
	ErrFormNotFound     = errors.New("Form Does Not Exist")
)

type Store interface {
	Create(response *model.FormResponse) error
	GetByID(responseID uint) (*model.FormResponse, error)
	GetByFormID(formID uint) ([]*model.FormResponse, error)
}
