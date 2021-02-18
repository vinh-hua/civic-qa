package models

// LogStore is an interface for storing and querying logs
type LogStore interface {
	Log(newEntry LogEntry) *LogError
	Query(query LogQuery) ([]LogEntry, *QueryError)
}
