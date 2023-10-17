package user

import (
	"fmt"
	"github.com/yuitaso/sampleWebServer/env"
	"github.com/yuitaso/sampleWebServer/src/entity/user"
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

func Create(newUser user.User) error {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return err
	}

	// Create
	db.Create(&UserModel{Name: newUser.Name, Password: newUser.Password})
	return nil
}

func FindById(id int) (user.User, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		fmt.Println("DB開くところでエラー")
	}

	var result UserModel
	if exec := db.First(&result, id); exec.Error != nil {
		return user.User{}, exec.Error
	}

	return user.User{Name: result.Name, Password: result.Password}, nil
}
