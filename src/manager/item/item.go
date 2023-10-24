package item

import (
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
	"gorm.io/gorm"
)

type ItemTable struct {
	gorm.Model
	gorm.DeletedAt
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

	if executed := manager.DB.Create(&data); executed.Error != nil {
		return nil, executed.Error
	}
	return &uuid, nil
}

func Update(item *entity.Item) error {
	values := generateUpdateValus(item)

	if executed := manager.DB.Model(&ItemTable{}).Where("uuid = ? ", item.Uuid.String()).Updates(values); executed.Error != nil {
		return executed.Error
	}
	return nil
}

func generateUpdateValus(item *entity.Item) *map[string]interface{} {
	val := map[string]interface{}{
		"price": item.Price,
	}
	return &val
}

func FindByUuid(id uuid.UUID) (*entity.Item, error) {
	var result ItemTable
	if executed := manager.DB.Where("uuid = ?", id.String()).First(&result); executed.Error != nil {
		return nil, executed.Error
	}

	return &entity.Item{
		Id:     result.ID,
		Uuid:   uuid.MustParse(result.Uuid),
		Price:  result.Price,
		UserId: result.UserId,
	}, nil
}

func Delete(id uuid.UUID) error {
	if executed := manager.DB.Where("uuid = ?", id.String()).Delete(&ItemTable{}); executed.Error != nil {
		return executed.Error
	}

	return nil
}
