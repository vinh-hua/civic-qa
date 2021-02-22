package model

import (
	"github.com/google/uuid"
)

// LogEntry represents a logged event
type LogEntry struct {
	ID            uint      `gorm:"primarykey"`
	CorrelationID uuid.UUID `gorm:"column:correlationID;type:string" json:"correlationID"`
	TimeUnix      int64     `gorm:"column:timeUnix" json:"timeUnix"`
	Service       string    `gorm:"column:service" json:"service"`
	StatusCode    int       `gorm:"column:statusCode" json:"statusCode"`
	Notes         string    `gorm:"column:notes" json:"notes"`
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
