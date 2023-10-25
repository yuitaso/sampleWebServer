package contract

import (
	"time"

	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
)

type ContractTable struct {
	ID               uint `gorm:"primarykey"`
	ItemId           uint
	CreatedAt        time.Time
	BuyerUserId      uint
	BuyerPointLogId  uint
	SellerPointLogId uint
}

func (c ContractTable) TableName() string {
	return "contract"
}

func Insert(user *entity.User, item *entity.Item, sellerPointLog *entity.PointLog, buyerPointLog *entity.PointLog) (*entity.Contract, error) {
	data := &ContractTable{
		BuyerUserId:      user.Id,
		ItemId:           item.Id,
		BuyerPointLogId:  buyerPointLog.Id,
		SellerPointLogId: sellerPointLog.Id,
	}

	if executed := manager.DB.Create(data); executed.Error != nil {
		return &entity.Contract{}, executed.Error
	}
	return &entity.Contract{
		ContractId:  data.ID,
		BuyerUserId: data.BuyerUserId,
		ItemId:      data.ItemId,
	}, nil
}
