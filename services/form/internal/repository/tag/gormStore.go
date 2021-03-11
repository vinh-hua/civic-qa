package tag

import (
	common "github.com/vivian-hua/civic-qa/services/common/model"

	"github.com/vivian-hua/civic-qa/services/form/internal/model"
	"gorm.io/gorm"
)

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
	err = db.AutoMigrate(&common.User{}, &model.Form{}, &model.FormResponse{}, &model.Tag{})
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

func (g *GormStore) GetAll(userID uint) ([]*model.Tag, error) {
	tags := make([]*model.Tag, 0)

	// find all responses
	result := g.db.
		Model(&model.Tag{}).
		Joins("JOIN formResponses ON formResponses.id = tags.formResponseID").
		Joins("JOIN forms ON forms.id = formResponses.formID").
		Where("forms.userID = ?", userID).
		Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	// return responses
	return tags, nil
}

func (g *GormStore) GetByResponseID(userID uint, responseID uint) ([]*model.Tag, error) {
	tags := make([]*model.Tag, 0)

	// find all responses to responseID
	result := g.db.
		Model(&model.Tag{}).
		Joins("JOIN formResponses ON formResponses.id = tags.formResponseID").
		Joins("JOIN forms ON forms.id = formResponses.formID").
		Where("forms.userID = ?", userID).
		Where("formResponses.id = ?", responseID).
		Find(&tags)

	if result.Error != nil {
		return nil, result.Error
	}

	// return responses
	return tags, nil
}

func (g *GormStore) Create(userID uint, responseID uint, value string) error {
	// find the associated response
	var response model.FormResponse
	result := g.db.
		Model(&model.FormResponse{}).
		Joins("JOIN forms ON forms.id = formResponses.formID").
		Where("formResponses.id = ?", responseID).
		Where("forms.userID = ?", userID).
		Take(&response)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrTagNotFound
		}
		return result.Error
	}

	// create the new tag
	tag := model.Tag{
		Value:          value,
		FormResponseID: response.ID,
	}

	result = g.db.Create(&tag)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (g *GormStore) Delete(userID uint, responseID uint, value string) error {
	// get the tag
	var tag model.Tag
	result := g.db.
		Model(&model.Tag{}).
		Joins("JOIN formResponses ON formResponses.id = tags.formResponseID").
		Joins("JOIN forms ON forms.id = formResponses.formID").
		Where("forms.userID = ?", userID).
		Where("formResponses.id = ?", responseID).
		Where("tags.value = ?", value).
		Take(&tag)

	// check for errors
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return ErrTagNotFound
		}
		return result.Error
	}

	// perform the delete
	result = g.db.Delete(&tag)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
