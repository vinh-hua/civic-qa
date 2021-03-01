package model

import "time"

const (
	// TableName used by TableName() to set SQL table name by Gorm
	TableName = "Users"
)

// User represents a user of CivicQA
type User struct {
	ID        uint      `gorm:"primarykey;column:id" json:"id"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	PassHash  []byte    `gorm:"column:passHash;not null" json:"-"`
	FirstName string    `gorm:"column:firstName;not null" json:"firstName"`
	LastName  string    `gorm:"column:lastName;not null" json:"lastName"`
	CreatedOn time.Time `gorm:"column:createdOn" json:"createdOn"`
}

// TableName implements Tabler interface
// https://gorm.io/docs/conventions.html#TableName
func (User) TableName() string {
	return TableName
}