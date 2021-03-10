package repository

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/vivian-hua/civic-qa/services/logAggregator/internal/model"
	"gorm.io/gorm"
)

// LogRepository interacts with log storage
type LogRepository struct {
	db *gorm.DB
}

// NewLogRepository returns a Log Repository for a given gorm Dialector
// returns an error if the connection cannot be established.
// performs initialization/migration for LogEntry model
func NewLogRepository(dialector gorm.Dialector, config *gorm.Config) (*LogRepository, error) {
	// Connect to the DB
	db, err := gorm.Open(dialector, config)
	if err != nil {
		return nil, err
	}

	// Initialize/migrate schema to match model
	db.AutoMigrate(model.LogEntry{})
	return &LogRepository{db: db}, nil
}

// Log stores a new LogEntry in the LogRepository
func (l *LogRepository) Log(newEntry model.LogEntry) *LogError {
	result := l.db.Create(&newEntry)
	if result.Error != nil {
		log.Printf("Error logging: %v", result.Error)
		return &LogError{errors.New("Error logging"), http.StatusInternalServerError}
	}

	return nil
}

// Query finds LogEntries that match a query
func (l *LogRepository) Query(query model.LogQuery) ([]model.LogEntry, *QueryError) {
	// get conditions and parameters for query
	conds, params := l.generateWheres(query)

	// build the query
	sess := l.db.Session(&gorm.Session{})
	for i := range conds {
		// apply each cond and parameter as a where clause
		sess = sess.Where(conds[i], params[i])
	}

	var entries []model.LogEntry

	// execute the query
	result := sess.Find(&entries).Order("timeUnix")

	if result.Error != nil {
		log.Printf("Error querying: %v", result.Error)
		return nil, &QueryError{errors.New("Error logging"), http.StatusInternalServerError}
	}

	return entries, nil
}

// generateWheres generates conditions and parameters for where clauses
// to query the LogRepository based on a LogQuery
func (l *LogRepository) generateWheres(query model.LogQuery) (conds, params []interface{}) {
	conds = make([]interface{}, 0)
	params = make([]interface{}, 0)

	// add a condition and parameter
	// for all non-default/zero fields
	if query.CorrelationID != uuid.Nil {
		conds = append(conds, "correlationID = ?")
		params = append(params, query.CorrelationID)
	}
	if query.TimeUnixStart != 0 {
		conds = append(conds, "timeUnix >= ?")
		params = append(params, query.TimeUnixStart)
	}
	if query.TimeUnixStop != 0 {
		conds = append(conds, "timeUnix <= ?")
		params = append(params, query.TimeUnixStop)
	}
	if query.HTTPMethod != "" {
		conds = append(conds, "httpMethod = ?")
		params = append(params, query.HTTPMethod)
	}
	if query.Service != "" {
		conds = append(conds, "service = ?")
		params = append(params, query.Service)
	}
	if query.StatusCodeStart != 0 {
		conds = append(conds, "statusCode >= ?")
		params = append(params, query.StatusCodeStart)
	}
	if query.StatusCodeStop != 0 {
		conds = append(conds, "statusCode <= ?")
		params = append(params, query.StatusCodeStop)
	}
	if query.Hostname != "" {
		conds = append(conds, "hostname = ?")
		params = append(params, query.Hostname)
	}

	return conds, params
}

// LogError represents an error creating a log
type LogError struct {
	Err  error
	Code int
}

// QueryError represents and error querying logs
type QueryError struct {
	Err  error
	Code int
}
