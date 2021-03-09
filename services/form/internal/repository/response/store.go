package response

import (
	"errors"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
)

// Query is used to query the response.Store
type Query struct {
	Name         string
	EmailAddress string
	InquiryType  string
	Subject      string
	ActiveOnly   bool
	TodayOnly    bool
	FormID       uint
}

var (
	// ErrResponseNotFound is returned when a requested FormResponse does not exist
	ErrResponseNotFound = errors.New("Response Does Not Exist")
)

// Store is an interface for implementations of response.Store.
// Describes storage of FormResponses
type Store interface {
	Create(response *model.FormResponse) error
	GetByID(responseID uint) (*model.FormResponse, error)
	PatchByID(userID uint, responseID uint, state bool) error
	GetResponses(userID uint, query Query) ([]*model.FormResponse, error)
}
