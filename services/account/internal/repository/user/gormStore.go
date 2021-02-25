package user

import (
	"github.com/vivian-hua/civic-qa/services/common/model"
	"gorm.io/gorm"
)

// GormStore implements user.Store
type GormStore struct {
	db *gorm.DB
}

// NewGormStore returns a GormStore based on a given gorm Dialector and gorm Config
// returns an error if the connection could not be opened or the migration fails
func NewGormStore(dialector gorm.Dialector, config *gorm.Config) (*GormStore, error) {
	// Open database with gorm
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	// perform schema migration
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return nil, err
	}

	// return the GormUserStore
	return &GormStore{db}, nil
}

// Create stores a new user, returns an error
// if the user cannot be created
func (g *GormStore) Create(user *model.User) error {
	result := g.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID returns a user from the store by their id
// returns an error if they cannot be found
func (g *GormStore) GetByID(id uint) (*model.User, error) {
	var user model.User
	result := g.db.Take(&user, id)
	if result.Error != nil {
		// return user not found
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		// or some unknown error
		return nil, result.Error
	}
	return &user, nil
}

// GetByEmail returns a user from the store by their email
// returns an error if they cannot be found
func (g *GormStore) GetByEmail(email string) (*model.User, error) {
	var user model.User
	result := g.db.Take("email = ?", email).First(&user)
	if result.Error != nil {
		// return user not found
		if result.Error == gorm.ErrRecordNotFound {
			return nil, ErrUserNotFound
		}
		// or some unknown error
		return nil, result.Error
	}
	return &user, nil
}
