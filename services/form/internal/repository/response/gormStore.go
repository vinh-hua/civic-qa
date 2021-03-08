package response

import (
	common "github.com/vivian-hua/civic-qa/services/common/model"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
	"gorm.io/gorm"
)

// GormStore implements response.Store
type GormStore struct {
	db *gorm.DB
}

// NewGormStore returns a GormStore based on a given gorm.Dialector, config, and a list
// of additional statements to execute after migration (useful for sqlite PRAGMAS or testing)
func NewGormStore(dialector gorm.Dialector, config *gorm.Config, exec ...string) (*GormStore, error) {
	// Open database with gorm
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	// perform schema migration
	err = db.AutoMigrate(&model.Form{}, &model.FormResponse{}, &common.User{})
	if err != nil {
		return nil, err
	}

	// perform execs
	for _, stmt := range exec {
		res := db.Exec(stmt)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	// return the GormUserStore
	return &GormStore{db}, nil
}

// Create stores a new FormResponse
func (g *GormStore) Create(response *model.FormResponse) error {
	result := g.db.Create(response)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID retrieves a new FormResponse by its ID
func (g *GormStore) GetByID(responseID uint) (*model.FormResponse, error) {
	var response model.FormResponse
	result := g.db.Take(&response, responseID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrResponseNotFound
		}
		return nil, result.Error
	}
	return &response, nil

}

// PatchByID updates the 'active' state of a FormResponse by its ID
func (g *GormStore) PatchByID(responseID uint, state bool) error {

	result := g.db.
		Model(&model.FormResponse{}).
		Where("id = ?", responseID).
		Update("active", state)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrResponseNotFound
		}
		return result.Error
	}

	return nil
}

// GetResponses returns a slice of FormResponses given a userID and a ResponseQuery.
// non-default fields in the ResponseQuery will be used to filter the returned FormResponses
func (g *GormStore) GetResponses(userID uint, query Query) ([]*model.FormResponse, error) {
	responses := make([]*model.FormResponse, 0)

	// apply non-default query parameters
	sess := applyQuery(query, g.db.Session(&gorm.Session{}))

	// execute the query
	result := sess.
		Joins("JOIN forms ON forms.ID = formResponses.formID").Where("forms.userID = ?", userID).
		Order("formResponses.createdAt DESC").
		Find(&responses)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrResponseNotFound
		}
		return nil, result.Error
	}

	// return the responses
	return responses, nil
}

func applyQuery(query Query, session *gorm.DB) *gorm.DB {
	if query.ActiveOnly {
		session = session.Where("formResponses.active = ?", query.ActiveOnly)
	}
	if query.EmailAddress != "" {
		session = session.Where("formResponses.emailAddress = ?", query.EmailAddress)
	}
	if query.FormID != 0 {
		session = session.Where("formResponses.formID = ?", query.FormID)
	}
	if query.Name != "" {
		session = session.Where("formResponses.name = ?", query.Name)
	}
	if query.Subject != "" {
		session = session.Where("formResponses.subject = ?", query.Subject)
	}

	return session
}
