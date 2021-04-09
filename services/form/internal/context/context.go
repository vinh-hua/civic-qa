package context

import (
	"fmt"

	"github.com/team-ravl/civic-qa/services/common/config"
	"github.com/team-ravl/civic-qa/services/form/internal/analytics"
	"github.com/team-ravl/civic-qa/services/form/internal/repository/form"
	"github.com/team-ravl/civic-qa/services/form/internal/repository/response"
	"github.com/team-ravl/civic-qa/services/form/internal/repository/tag"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Context stores request handler context
type Context struct {
	FormStore     form.Store
	ResponseStore response.Store
	TagStore      tag.Store
	Analytics     analytics.Client
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

	tagStore, err := getTagStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	analyticsClient, err := getAnalyticsClient(cfg)
	if err != nil {
		return nil, err
	}

	return &Context{
		FormStore:     formStore,
		ResponseStore: respStore,
		TagStore:      tagStore,
		Analytics:     analyticsClient,
	}, nil
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
	case "mysql":
		formStore, err := form.NewGormStore(mysql.Open(dbDsn), &gorm.Config{})
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
	case "mysql":
		respStore, err := response.NewGormStore(mysql.Open(dbDsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return respStore, nil
	default:
		return nil, fmt.Errorf("Unknown DB_IMPL: %s", dbImpl)
	}
}

// getTagStoreImpl returns a tag.Store implementation based on
// a given config Provider
func getTagStoreImpl(cfg config.Provider) (tag.Store, error) {
	dbImpl := cfg.GetOrFallback("DB_IMPL", "sqlite")
	dbDsn := cfg.GetOrFallback("DB_DSN", "database.db")
	switch dbImpl {
	case "sqlite":
		tagStore, err := tag.NewGormStore(sqlite.Open(dbDsn), &gorm.Config{}, "PRAGMA foreign_keys = ON;")
		if err != nil {
			return nil, err
		}
		return tagStore, nil
	case "mysql":
		tagStore, err := tag.NewGormStore(mysql.Open(dbDsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return tagStore, nil
	default:
		return nil, fmt.Errorf("Unknown DB_IMPL: %s", dbImpl)
	}
}

func getAnalyticsClient(cfg config.Provider) (analytics.Client, error) {
	analyticsImpl := cfg.GetOrFallback("ANALYTICS_IMPL", "v0")
	switch analyticsImpl {
	case "v0":
		analyticsAddr := cfg.GetOrFallback("ANALYTICS_ADDR", "http://localhost:9090")
		analyticsClient, err := analytics.NewClientV0(analyticsAddr)
		if err != nil {
			return nil, err
		}
		return analyticsClient, nil
	default:
		return nil, fmt.Errorf("Unknown ANALYTICS_IMPL: %s", analyticsImpl)
	}
}
