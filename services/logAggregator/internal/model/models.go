package model

import (
	"github.com/google/uuid"
)

// LogEntry represents a logged event
type LogEntry struct {
	ID            uint      `gorm:"primarykey" json:"-"`
	CorrelationID uuid.UUID `gorm:"column:correlationID;type:string" json:"correlationID"`
	TimeUnix      int64     `gorm:"column:timeUnix" json:"timeUnix"`
	HTTPMethod    string    `gorm:"column:httpMethod" json:"httpMethod"`
	RequestPath   string    `gorm:"column:requestPath" json:"requestPath"`
	Service       string    `gorm:"column:service" json:"service"`
	StatusCode    int       `gorm:"column:statusCode" json:"statusCode"`
	Hostname      string    `gorm:"column:hostname" json:"hostname"`
	Notes         string    `gorm:"column:notes" json:"notes"`
}

// TableName overrides default gorm table naming
func (LogEntry) TableName() string {
	return "logs"
}

// LogQuery represents a query against existing logs
type LogQuery struct {
	CorrelationID   uuid.UUID `json:"correlationID"`
	TimeUnixStart   int64     `json:"timeUnixStart"`
	TimeUnixStop    int64     `json:"timeUnixStop"`
	HTTPMethod      string    `json:"httpMethod"`
	Service         string    `json:"service"`
	StatusCodeStart int       `json:"statusCodeStart"`
	StatusCodeStop  int       `json:"statusCodeStop"`
	Hostname        string    `json:"hostname"`
}
