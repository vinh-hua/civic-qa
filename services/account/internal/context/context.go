package context

import (
	"fmt"

	"github.com/vivian-hua/civic-qa/service/account/internal/repository/session"
	"github.com/vivian-hua/civic-qa/service/account/internal/repository/user"
	"github.com/vivian-hua/civic-qa/services/common/environment"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Context contains handler context information
type Context struct {
	SessionStore session.Store
	UserStore    user.Store
}

// BuildContext creates a the handler context
// for requests using environment variables,
// returns an error if the context could not be created
func BuildContext() (*Context, error) {
	var sessionStore session.Store
	var userStore user.Store

	// Initialize SessionStore based on environment variables
	sessImpl := environment.GetEnvOrFallback("SESS_IMPL", "memory")
	switch sessImpl {
	case "memory":
		sessionStore = session.NewMemoryStore()
	default:
		return nil, fmt.Errorf("Unknown SESS_IMPL: %s", sessImpl)
	}

	// Initialize UserStore based on environment variables
	dbImpl := environment.GetEnvOrFallback("DB_IMPL", "sqlite")
	dbDsn := environment.GetEnvOrFallback("DB_DSN", "accounts.db")
	switch dbImpl {
	case "sqlite":
		store, err := user.NewGormStore(sqlite.Open(dbDsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}

		userStore = store
	default:
		return nil, fmt.Errorf("UNKNOWN DB_IMPL: %s", dbImpl)
	}

	return &Context{SessionStore: sessionStore, UserStore: userStore}, nil
}
