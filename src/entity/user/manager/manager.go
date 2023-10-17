package manager

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

func FindById(id int) user.User {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		// err
		fmt.Println("DB開くところでエラー")
	}

	var res UserModel
	db.First(&res, id)
	fmt.Println("めも")
	fmt.Println(res)

	return user.User{Name: res.Name, Password: res.Password}
}
