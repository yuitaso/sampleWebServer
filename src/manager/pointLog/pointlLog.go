package pointlog

import (
	"fmt"
	"time"

	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
)

type PointLog struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index:idx_user"`
	UserId    uint
	Amount    int
}

func Insert(user *entity.User, amount int) error {
	data := &PointLog{UserId: user.Id, Amount: amount}

	if executed := manager.DB.Create(data); executed.Error != nil {
		return executed.Error
	}
	return nil
}

func GetSum(user *entity.User) (int, error) {
	var result *interface{}
	if executed := manager.DB.Model(&PointLog{}).
		Select("sum(amount) as total").Where("user_id = ?", user.Id).Group("name").First(result); executed.Error != nil {
		return 0, nil
	}

	fmt.Println(result)
	return 1, nil
}
