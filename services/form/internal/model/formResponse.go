package model

type FormResponse struct {
	ID           uint   `gorm:"primarykey;column:id" json:"id"`
	EmailAddress string `gorm:"column:emailAddress" json:"emailAddress"`
	Subject      string `gorm:"column:subject" json:"subject"`
	Body         string `gorm:"column:body" json:"body"`
	Form         Form   `gorm:"column:formID" json:"formID"`
}
