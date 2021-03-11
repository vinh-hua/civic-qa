package model

import (
	"time"

	common "github.com/team-ravl/civic-qa/services/common/model"
)

const (
	// FormTableName used by TableName() to set SQL table name by Gorm
	FormTableName = "forms"
	// FormResponseTableName used by TableName() to set SQL table name by Gorm
	FormResponseTableName = "formResponses"
	// TagTableName used by TableName() to set SQL table name by Gorm
	TagTableName = "tags"
)

// Form is a model for a generated HTML form
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

// FormResponse is a model for a response to an HTML form
type FormResponse struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	Name         string    `gorm:"column:name" json:"name"`
	EmailAddress string    `gorm:"column:emailAddress" json:"emailAddress"`
	InquiryType  string    `gorm:"column:inquiryType" json:"inquiryType"`
	Subject      string    `gorm:"column:subject" json:"subject"`
	Body         string    `gorm:"column:body" json:"body"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"createdAt"`
	Active       bool      `gorm:"column:active" json:"active"`
	FormID       uint      `gorm:"column:formID" json:"formID"`
	Form         Form      `gorm:"foreignKey:FormID" json:"-"`
	Tags         []Tag     `json:"tags"`
}

// TableName implements Tabler interface
func (FormResponse) TableName() string {
	return FormResponseTableName
}

// Tag is a model for a tag on a FormResponse
type Tag struct {
	ID             uint   `gorm:"primaryKey;column:id" json:"-"`
	Value          string `gorm:"index:idx_value;column:value" json:"value"`
	FormResponseID uint   `gorm:"column:formResponseID" json:"-"`
}

// TableName implements Tabler interface
func (Tag) TableName() string {
	return TagTableName
}
