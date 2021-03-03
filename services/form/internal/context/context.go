package context

import (
	"fmt"

	"github.com/vivian-hua/civic-qa/services/common/config"
	"github.com/vivian-hua/civic-qa/services/form/internal/repository/form"
	"github.com/vivian-hua/civic-qa/services/form/internal/repository/response"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Context stores request handler context
type Context struct {
	FormStore     form.Store
	ResponseStore response.Store
}

// BuildContext returns a handler Context given a config Provider
func BuildContext(cfg config.Provider) (*Context, error) {
	formStore, err := getFormStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	respStore, err := getResponseStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	return &Context{FormStore: formStore, ResponseStore: respStore}, nil
}

// getFormStoreImpl returns a form.Store implementation based on
// a given config Provider
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

// getResponseStoreImpl returns a response.Store implementation based on
// a given config Provider
func getResponseStoreImpl(cfg config.Provider) (response.Store, error) {
	dbImpl := cfg.GetOrFallback("DB_IMPL", "sqlite")
	dbDsn := cfg.GetOrFallback("DB_DSN", "database.db")
	switch dbImpl {
	case "sqlite":
		respStore, err := response.NewGormStore(sqlite.Open(dbDsn), &gorm.Config{}, "PRAGMA foreign_keys = ON;")
		if err != nil {
			return nil, err
		}
		return respStore, nil
	default:
		return nil, fmt.Errorf("Unknown DB_IMPL: %s", dbImpl)
	}
}
