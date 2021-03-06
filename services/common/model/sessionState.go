package model

import "time"

// SessionState represents a user session
type SessionState struct {
	User      User      `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
}
