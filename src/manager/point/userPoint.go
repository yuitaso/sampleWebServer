package point

import (
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
	"gorm.io/gorm"
)

type UserPointTable struct {
	gorm.Model
	UserID uint
	Amount int
}

func Insert(user *entity.User, amount int) (int, error) {
	data := UserPointTable{UserID: user.Id, Amount: amount}
	if executed := manager.DB.Create(&data); executed.Error != nil {
		return 0, executed.Error
	}
	return amount, nil
}

func Update(user *entity.User, amount int) (int, error) {
	data := UserPointTable{UserID: user.Id, Amount: amount}
	if executed := manager.DB.Create(&data); executed.Error != nil {
		return 0, executed.Error
	}
	return amount, nil
}
