package model

import (
	common "github.com/vivian-hua/civic-qa/services/common/model"
)

type Form struct {
	// TODO: overiding the colum name for userID seems to cause issues, test more
	ID     uint        `gorm:"primarykey;column:id" json:"id"`
	UserID uint        `gorm:"column:userID" json:"userID"`
	User   common.User `json:"-"` // Gorm belongs to (https://gorm.io/docs/belongs_to.html)
}
