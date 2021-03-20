package context

import (
	"errors"

	"github.com/team-ravl/civic-qa/services/common/config"
	"github.com/team-ravl/civic-qa/services/logAggregator/internal/repository"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Context stores handler context information for logAggregator
type Context struct {
	Repo *repository.LogRepository
}

func BuildContext(cfg config.Provider) (*Context, error) {
	logRepo, err := getLogRepositoryImpl(cfg)
	if err != nil {
		return nil, err
	}
	return &Context{Repo: logRepo}, nil
}

func getLogRepositoryImpl(cfg config.Provider) (*repository.LogRepository, error) {
	dbImpl := cfg.GetOrFallback("DB_IMPL", "sqlite")
	dbDsn := cfg.GetOrFallback("DB_DSN_LOGS", "logs.db")
	switch dbImpl {
	case "sqlite":
		logRepo, err := repository.NewLogRepository(sqlite.Open(dbDsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return logRepo, nil
	case "mysql":
		logRepo, err := repository.NewLogRepository(mysql.Open(dbDsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return logRepo, nil
	default:
		return nil, errors.New("Unknown DB_IMPL")
	}
}
