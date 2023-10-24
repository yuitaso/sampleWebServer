package item

import (
	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/env"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

type ItemTable struct {
	gorm.Model
	Uuid      string
	Price     int
	UserId    uint
	DeletedAt soft_delete.DeletedAt
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

func Update(item *entity.Item) error {
	values := generateUpdateValus(item)
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return err
	}

	if executed := db.Model(&ItemTable{}).Where("uuid = ? ", item.Uuid.String()).Updates(values); executed.Error != nil {
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
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	var result ItemTable
	if executed := db.Where("uuid =  ?", id.String()).First(&result); executed.Error != nil {
		return nil, executed.Error
	}

	return &entity.Item{Id: result.ID, Uuid: uuid.MustParse(result.Uuid), Price: result.Price, UserId: result.UserId}, nil
}

func Delete(id uuid.UUID) error {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return err
	}

	if executed := db.Where("uuid = ?", id.String()).Delete(&ItemTable{}); executed.Error != nil {
		return executed.Error
	}

	return nil
}
