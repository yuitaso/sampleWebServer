package item

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	contractManager "github.com/yuitaso/sampleWebServer/src/manager/contract"
	pointLogManager "github.com/yuitaso/sampleWebServer/src/manager/pointLog"
	userManager "github.com/yuitaso/sampleWebServer/src/manager/user"
)

type ItemTable struct {
	gorm.Model
	gorm.DeletedAt
	Uuid   string
	Price  int
	UserId uint
	Sold   bool // soldedAtでもいいかも
}

func (i ItemTable) TableName() string {
	return "item"
}

func Insert(user *entity.User, price int) (*uuid.UUID, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	data := ItemTable{Uuid: uuid.String(), Price: price, UserId: user.Id, Sold: false}

	if executed := manager.DB.Create(&data); executed.Error != nil {
		return nil, executed.Error
	}
	return &uuid, nil
}

func Update(item *entity.Item) error {
	data := map[string]interface{}{
		"price": item.Price,
	}
	if executed := manager.DB.Model(&ItemTable{}).Where("uuid = ? ", item.Uuid.String()).Updates(&data); executed.Error != nil {
		return executed.Error
	}
	return nil
}

func FindByUuid(id uuid.UUID) (*entity.Item, error) {
	var result ItemTable
	if executed := manager.DB.Where("uuid = ?", id.String()).First(&result); executed.Error != nil {
		return nil, executed.Error
	}

	return &entity.Item{
		Id:           result.ID,
		Uuid:         uuid.MustParse(result.Uuid),
		Price:        result.Price,
		SellerUserId: result.UserId,
	}, nil
}

func FindById(id uint) (*entity.Item, error) {
	var result ItemTable
	if executed := manager.DB.Where("id = ?", id).First(&result); executed.Error != nil {
		return nil, executed.Error
	}

	return &entity.Item{
		Id:           result.ID,
		Uuid:         uuid.MustParse(result.Uuid),
		Price:        result.Price,
		SellerUserId: result.UserId,
	}, nil
}

func DeleteByUuid(item_uuid uuid.UUID) error {
	if executed := manager.DB.Where("uuid = ?", item_uuid.String()).Delete(&ItemTable{}); executed.Error != nil {
		return executed.Error
	}

	return nil
}

func SetSold(item *entity.Item) error {
	data := map[string]interface{}{
		"sold": false,
	}
	if executed := manager.DB.Model(&ItemTable{}).Where("uuid = ? ", item.Uuid.String()).Updates(&data); executed.Error != nil {
		return executed.Error
	}
	return nil
}

func Buy(buyer *entity.User, item *entity.Item) error {
	// validate
	hasPoint, err := pointLogManager.FetchCurrentPoint(buyer)
	if err != nil {
		return err
	}
	if hasPoint < item.Price {
		return errors.New("You do not have enough point.")
	}

	err = manager.DB.Transaction(func(db *gorm.DB) error {
		// ロック取る
		var lock map[string]interface{}
		manager.DB.Model(&ItemTable{}).Where("uuid = ? ", item.Uuid.String()).Clauses(clause.Locking{Strength: "UPDATE"}).Find(&lock)
		fmt.Println(lock)

		// アイテムを売り切れに
		if err := SetSold(item); err != nil {
			return err
		}

		// Pointを精算
		buyerLog, err := pointLogManager.UsePoint(buyer, item.Price)
		if err != nil {
			return err
		}
		seller, err := userManager.FindById(item.SellerUserId)
		sellerLog, err := pointLogManager.GrantPoint(seller, item.Price)
		if err != nil {
			return err
		}

		// Contractを作成
		_, err = contractManager.Insert(buyer, item, sellerLog, buyerLog)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
