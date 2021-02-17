package models

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	// sqlite drivers
	_ "github.com/mattn/go-sqlite3"
)

const (
	driverName = "sqlite3"
)

// SQLiteLogStore is a SQLite backed implementation
// of LogStore
type SQLiteLogStore struct {
	LogStore
	*sql.DB
}

// NewSQLiteLogStore retuns a new SQLiteLogStore
// for a given dsn.
// if initDB, the database schema will be initialized.
// Logs fatally if fails.
func NewSQLiteLogStore(dsn string, initDB bool) *SQLiteLogStore {
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalf("Failed to open SQLite DB: %v", err)
	}
	store := &SQLiteLogStore{DB: db}
	if initDB {
		store.initializeDB()
	}

	return store
}

// Log stores a new LogEntry in the database
func (s *SQLiteLogStore) Log(newEntry LogEntry) *LogError {
	// prepare insert
	ins, err := s.DB.Prepare("INSERT INTO Logs (correlationID, timeUnix, service, statusCode, notes) VALUES (?,?,?,?,?)")
	if err != nil {
		return &LogError{Err: err, Code: 500}
	}

	// execute insert
	_, err = ins.Exec(
		newEntry.CorrelationID.String(),
		newEntry.TimeUnix,
		newEntry.Service,
		newEntry.StatusCode,
		newEntry.Notes,
	)
	if err != nil {
		return &LogError{Err: err, Code: 500}
	}

	return nil
}

// Query returns LogEntries matching a LogQuery.
// TODO: UNIMPLEMENTED
func (s *SQLiteLogStore) Query(query LogQuery) ([]*LogEntry, *QueryError) {
	return nil, &QueryError{Err: errors.New("Unimplemented"), Code: http.StatusNotImplemented}
}

// Initializes the Users table in SQLite syntax, for testing mostly
func (s *SQLiteLogStore) initializeDB() {

	log.Println("Initializing database...")
	stmt, err := s.DB.Prepare(`
		CREATE TABLE IF NOT EXISTS Logs (
			logID INTEGER PRIMARY KEY AUTOINCREMENT,
			correlationID TEXT NOT NULL,
			timeUnix INTEGER NOT NULL,
			service TEXT NOT NULL,
			statusCode INTEGER NOT NULL,
			notes TEXT
		)`)

	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatalf("Error initializing DB: %v", err)
	}
	log.Println("Database initialized!")
}
