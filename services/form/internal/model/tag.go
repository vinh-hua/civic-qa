package model

const (
	// TagTableName used by TableName() to set SQL table name by Gorm
	TagTableName = "tags"
)

// Tag is a model for a tag on a FormResponse
type Tag struct {
	ID             uint         `gorm:"primaryKey;column:id" json:"-"`
	Value          string       `gorm:"index:idx_value;column:value" json:"value"`
	FormResponseID uint         `gorm:"column:formResponseID" json:"-"`
	FormResponse   FormResponse `gorm:"foreignKey:FormResponseID" json:"-"`
}

// TableName implements Tabler interface
func (Tag) TableName() string {
	return TagTableName
}
