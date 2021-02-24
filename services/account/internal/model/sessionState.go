package model

import (
	"time"

	common "github.com/vivian-hua/civic-qa/services/common/model"
)

type SessionState struct {
	User      common.User
	CreatedAt time.Time `json:"createdAt"`
}
