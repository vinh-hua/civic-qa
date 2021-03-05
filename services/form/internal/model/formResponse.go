package model

import "time"

const (
	// FormResponseTableName used by TableName() to set SQL table name by Gorm
	FormResponseTableName = "formResponses"
)

// FormResponse is a model for a response to an HTML form
type FormResponse struct {
	ID           uint      `gorm:"primaryKey;column:id" json:"id"`
	EmailAddress string    `gorm:"column:emailAddress" json:"emailAddress"`
	Subject      string    `gorm:"column:subject" json:"subject"`
	Body         string    `gorm:"column:body" json:"body"`
	CreatedAt    time.Time `gorm:"column:createdAt" json:"createdAt"`
	Open         bool      `gorm:"column:open" json:"open"`
	FormID       uint      `gorm:"column:formID" json:"formID"`
	Form         Form      `gorm:"foreignKey:FormID" json:"-"`
}

// TableName implements Tabler interface
func (FormResponse) TableName() string {
	return FormResponseTableName
}
