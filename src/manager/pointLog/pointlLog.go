package pointlog

import (
	"time"

	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
)

type PointLogTable struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UserId    uint `gorm:"index:idx_user"`
	Amount    int
}

func (p PointLogTable) TableName() string {
	return "pointLog"
}

func UsePoint(user *entity.User, amount int) (*entity.PointLog, error) {
	return Insert(user, -1*amount)
}

func GrantPoint(user *entity.User, amount int) (*entity.PointLog, error) {
	return Insert(user, amount)
}

func Insert(user *entity.User, amount int) (*entity.PointLog, error) {
	data := &PointLogTable{UserId: user.Id, Amount: amount}

	if executed := manager.DB.Create(data); executed.Error != nil {
		return &entity.PointLog{}, executed.Error
	}
	return &entity.PointLog{
		Id:        data.ID,
		UserId:    data.UserId,
		Amount:    data.Amount,
		CreatedAt: data.CreatedAt,
	}, nil
}

func FetchCurrentPoint(user *entity.User) (int, error) {
	type result struct {
		TotalAmount int
	}
	var res result
	if executed := manager.DB.Model(&PointLogTable{}).Select("sum(amount) as total_amount").Where("user_id = ?", user.Id).First(&res); executed.Error != nil {
		return 0, executed.Error
	}

	return res.TotalAmount, nil
}
