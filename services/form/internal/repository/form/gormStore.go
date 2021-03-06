package form

import (
	common "github.com/vivian-hua/civic-qa/services/common/model"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
	"gorm.io/gorm"
)

// GormStore is a gorm implementation
// of form.Store
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
	err = db.AutoMigrate(&model.Form{}, &common.User{})
	if err != nil {
		return nil, err
	}

	// perform execs
	for _, stmt := range exec {
		err = db.Exec(stmt).Error
		if err != nil {
			return nil, err
		}
	}

	// return the GormUserStore
	return &GormStore{db}, nil
}

// Create stores a new Form
func (g *GormStore) Create(form *model.Form) error {
	result := g.db.Create(form)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetByID returns a form given its ID
func (g *GormStore) GetByID(formID uint) (*model.Form, error) {
	var form model.Form
	result := g.db.Take(&form, formID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrFormNotFound
		}
		return nil, result.Error
	}

	return &form, nil
}

// GetByUserID returns all Forms of a given userID
func (g *GormStore) GetByUserID(userID uint) ([]*model.Form, error) {
	forms := make([]*model.Form, 0)
	result := g.db.Where("userID = ?", userID).Find(&forms)
	if result.Error != nil {
		return nil, result.Error
	}

	return forms, nil
}
