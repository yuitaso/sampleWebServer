package point

import (
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserPointTable struct {
	gorm.Model
	UserID uint
	Amount int
}

func Insert(user *entity.User, amount int) (int, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return 0, err
	}

	data := UserPointTable{UserID: user.Id, Amount: amount}
	if executed := db.Create(&data); executed.Error != nil {
		return 0, executed.Error
	}
	return amount, nil
}

func Update(user *entity.User, amount int) (int, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return 0, err
	}

	data := UserPointTable{UserID: user.Id, Amount: amount}
	if executed := db.Create(&data); executed.Error != nil {
		return 0, executed.Error
	}
	return amount, nil
}
