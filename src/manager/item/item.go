package item

import "gorm.io/gorm"

type ItemTable struct {
	gorm.Model
	Uuid  string
	Price int
}

func (i ItemTable) TableName() string {
	return "user"
}
