package models

import (
	"github.com/google/uuid"
)

// LogEntry represents a logged event
type LogEntry struct {
	CorrelationID uuid.UUID `json:"correlationID"`
	TimeUnix      int64     `json:"timeUnix"`
	Service       string    `json:"service"`
	StatusCode    int       `json:"statusCode"`
	Notes         string    `json:"Notes"`
}

// LogQuery represents a query against existing logs
type LogQuery struct {
	CorrelationID   uuid.UUID `json:"correlationID"`
	TimeUnixStart   int64     `json:"timeUnixStart"`
	TimeUnixStop    int64     `json:"timeUnixStop"`
	Service         string    `json:"service"`
	StatusCodeStart int       `json:"statusCodeStart"`
	StatusCodeStop  int       `json:"statusCodeStop"`
}

type LogError struct {
	Err  error
	Code int
}

type QueryError struct {
	Err  error
	Code int
}
