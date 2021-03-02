package model

import (
	"time"

	common "github.com/vivian-hua/civic-qa/services/common/model"
)

const (
	// FormTableName used by TableName() to set SQL table name by Gorm
	FormTableName = "forms"
)

type Form struct {
	ID        uint        `gorm:"primarykey;column:id" json:"id"`
	Name      string      `gorm:"NOT NULL;column:name" json:"name"`
	CreatedAt time.Time   `gorm:"column:createdAt" json:"createdAt"`
	UserID    uint        `gorm:"column:userID" json:"userID"`
	User      common.User `gorm:"foreignKey:UserID" json:"-"` // Gorm belongs to (https://gorm.io/docs/belongs_to.html)
}

// TableName implements Tabler interface
func (Form) TableName() string {
	return FormTableName
}
