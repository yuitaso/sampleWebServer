package user

import (
	"fmt"
	"github.com/yuitaso/sampleWebServer/env"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserModel struct {
	gorm.Model
	Name     string
	Password string
}

func (u UserModel) TableName() string {
	return "user"
}

func Create(newUser entity.User) (uint, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return 0, err
	}

	model := UserModel{Name: newUser.Name, Password: newUser.Password}
	if executed := db.Create(&model); executed.Error != nil {
		return 0, executed.Error
	}
	return model.ID, nil
}

func FindById(id int) (entity.User, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		fmt.Println("DB開くところでエラー")
	}

	var result UserModel
	if executed := db.First(&result, id); executed.Error != nil {
		return entity.User{}, executed.Error
	}

	return entity.User{Name: result.Name, Password: result.Password}, nil
}
