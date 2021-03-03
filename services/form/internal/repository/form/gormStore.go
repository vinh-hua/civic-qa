package form

import (
	common "github.com/vivian-hua/civic-qa/services/common/model"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
	"gorm.io/gorm"
)

type GormStore struct {
	db *gorm.DB
}

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
		res := db.Exec(stmt)
		if res.Error != nil {
			return nil, res.Error
		}
	}

	// return the GormUserStore
	return &GormStore{db}, nil
}

func (g *GormStore) Create(form *model.Form) error {
	result := g.db.Create(form)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
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

func (g *GormStore) GetByUserID(userID uint) ([]*model.Form, error) {
	forms := make([]*model.Form, 0)
	result := g.db.Where("userID = ?", userID).Find(&forms)
	if result.Error != nil {
		return nil, result.Error
	}

	return forms, nil
}
