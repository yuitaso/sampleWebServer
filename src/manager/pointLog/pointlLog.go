package pointlog

import (
	"fmt"

	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
	"gorm.io/gorm"
)

type PointLog struct {
	gorm.Model
	ID uint `gorm:"primarykey"`
	// CreatedAt time.Time `gorm:"index:idx_user"`
	UserId uint
	Amount int
}

func (p PointLog) TableName() string {
	return "pointLog"
}

func Insert(user *entity.User, amount int) error {
	data := &PointLog{UserId: user.Id, Amount: amount}

	if executed := manager.DB.Create(data); executed.Error != nil {
		return executed.Error
	}
	return nil
}

func GetSum(user *entity.User) (int, error) {

	type result struct {
		TotalAmount int
	}
	var res result
	fmt.Println("ゆーざー", user.Id)
	// hoge := manager.DB.Debug().Raw("select sum(amount) as total from pointLog where user_id = ?", user.Id)
	// fmt.Println(hoge)
	if executed := manager.DB.Model(&PointLog{}).Select("sum(amount) as total_amount").Where("user_id = 1").First(&res); executed.Error != nil {
		fmt.Println("ほげ", executed)
		return 0, executed.Error
	}

	return res.TotalAmount, nil
}
