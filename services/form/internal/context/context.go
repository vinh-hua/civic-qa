package context

import (
	"fmt"

	"github.com/vivian-hua/civic-qa/services/common/config"
	"github.com/vivian-hua/civic-qa/services/form/internal/repository/form"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Context struct {
	FormStore form.Store
}

func BuildContext(cfg config.Provider) (*Context, error) {
	formStore, err := getFormStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	return &Context{FormStore: formStore}, nil
}

func getFormStoreImpl(cfg config.Provider) (form.Store, error) {
	dbImpl := cfg.GetOrFallback("DB_IMPL", "sqlite")
	dbDsn := cfg.GetOrFallback("DB_DSN", "database.db")
	switch dbImpl {
	case "sqlite":
		formStore, err := form.NewGormStore(sqlite.Open(dbDsn), &gorm.Config{}, "PRAGMA foreign_keys = ON;")
		if err != nil {
			return nil, err
		}
		return formStore, nil
	default:
		return nil, fmt.Errorf("Unknown DB_IMPL: %s", dbImpl)
	}

}
