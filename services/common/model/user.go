package model

import "time"

type User struct {
	ID        uint      `gorm:"primarykey;column:id" json:"id"`
	Email     string    `gorm:"column:email;unique;not null" json:"email"`
	PassHash  []byte    `gorm:"column:passHash;not null" json:"-"`
	FirstName string    `gorm:"not null" json:"firstName"`
	LastName  string    `gorm:"not null" json:"lastName"`
	CreatedOn time.Time `json:"createdOn"`
}
