package item

import (
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ItemTable struct {
	gorm.Model
	Uuid   string
	Price  int
	UserId uint
}

func (i ItemTable) TableName() string {
	return "item"
}

func Insert(user *entity.User, price int) (*uuid.UUID, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	data := ItemTable{Uuid: uuid.String(), Price: price, UserId: user.Id}
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if executed := db.Create(&data); executed.Error != nil {
		return nil, executed.Error
	}
	return &uuid, nil
}
