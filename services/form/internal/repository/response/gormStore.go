package response

import (
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
	err = db.AutoMigrate(&model.Form{}, &model.FormResponse{})
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

func (g *GormStore) Create(response *model.FormResponse) error {
	result := g.db.Create(response)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (g *GormStore) GetByID(responseID uint) (*model.FormResponse, error) {
	var response model.FormResponse
	result := g.db.Take(&response, responseID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrFormNotFound
		}
		return nil, result.Error
	}
	return &response, nil

}
func (g *GormStore) GetByFormID(formID uint) ([]*model.FormResponse, error) {
	responses := make([]*model.FormResponse, 0)
	result := g.db.Where("formID = ?", formID).Find(&responses)
	if result.Error != nil {
		return nil, result.Error
	}

	return responses, nil
}
