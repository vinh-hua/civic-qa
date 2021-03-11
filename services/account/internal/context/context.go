// Package context defines the Context struct, which holds handler context
// information for each request
package context

import (
	"fmt"

	"github.com/team-ravl/civic-qa/service/account/internal/repository/session"
	"github.com/team-ravl/civic-qa/service/account/internal/repository/user"
	"github.com/team-ravl/civic-qa/services/common/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Context contains handler context information
type Context struct {
	SessionStore session.Store
	UserStore    user.Store
}

// BuildContext creates the handler context
// for requests using environment variables,
// returns an error if the context could not be created
func BuildContext(cfg config.Provider) (*Context, error) {

	// get the session.Store implementation
	sessionStore, err := getSessionStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	// get the User.Store implementation
	userStore, err := getUserStoreImpl(cfg)
	if err != nil {
		return nil, err
	}

	return &Context{SessionStore: sessionStore, UserStore: userStore}, nil
}

// getSessionStoreImpl returns a session.Store given a config.Provider
// or returns an error if one cannot be created with the given config
func getSessionStoreImpl(cfg config.Provider) (session.Store, error) {
	// Initialize SessionStore based on environment variables
	sessImpl := cfg.GetOrFallback("SESS_IMPL", "memory")
	switch sessImpl {
	case "memory":
		return session.NewMemoryStore(), nil
	case "redis":
		redisAddr := cfg.GetOrFallback("REDIS_ADDR", "localhost:6379")
		return session.NewRedisStore(redisAddr)
	default:
		return nil, fmt.Errorf("Unknown SESS_IMPL: %s", sessImpl)
	}
}

// getUserStoreImpl returns a User.Store given a config.Provider
// or returns an error if one cannot be created with the given config
func getUserStoreImpl(cfg config.Provider) (user.Store, error) {
	// Initialize UserStore based on environment variables
	dbImpl := cfg.GetOrFallback("DB_IMPL", "sqlite")
	dbDsn := cfg.GetOrFallback("DB_DSN", "database.db")
	switch dbImpl {
	case "sqlite":
		store, err := user.NewGormStore(sqlite.Open(dbDsn), &gorm.Config{}, "PRAGMA foreign_keys = ON;")
		if err != nil {
			return nil, err
		}

		return store, nil
	default:
		return nil, fmt.Errorf("UNKNOWN DB_IMPL: %s", dbImpl)
	}
}
