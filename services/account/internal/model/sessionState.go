// Package model defines data models used by account
package model

import (
	"time"

	common "github.com/vivian-hua/civic-qa/services/common/model"
)

// SessionState represents a user session
type SessionState struct {
	User      common.User `json:"user"`
	CreatedAt time.Time   `json:"createdAt"`
}
