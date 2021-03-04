package response

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func makeStore() *GormStore {
	db := sqlite.Open(":memory:")
	gs, err := NewGormStore(db, &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gs
}
