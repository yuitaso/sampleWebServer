package user

import (
	"errors"
	"fmt"
	"github.com/yuitaso/sampleWebServer/env"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

type UserTable struct {
	gorm.Model
	Email    string
	Password string
}

func (u UserTable) TableName() string {
	return "user"
}

func Create(email string, rawPass string) (uint, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.MinCost) // fix me コスト最適化
	data := UserTable{Email: email, Password: string(hashed)}
	err = data.Validate()
	if err != nil {
		return 0, err
	}

	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return 0, err
	}

	if executed := db.Create(&data); executed.Error != nil {
		return 0, executed.Error
	}
	return data.ID, nil
}

func FindById(id int) (entity.User, error) {
	db, err := gorm.Open(sqlite.Open(env.DbName), &gorm.Config{})
	if err != nil {
		return entity.User{}, err
	}

	var result UserTable
	if executed := db.First(&result, id); executed.Error != nil {
		return entity.User{}, executed.Error
	}

	return entity.User{Email: result.Email}, nil // TODO Email
}

func Authenticate(email string, password string) error {
	var db *gorm.DB
	var err error = nil

	if db, err = gorm.Open(sqlite.Open(env.DbName), &gorm.Config{}); err != nil {
		return err
	}

	var result UserTable
	if executed := db.First(&result, "email = ?", email); executed.Error != nil {
		return executed.Error
	}

	return bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)) // imo 毎回キャストするとメモリ効率が悪いのでいい感じにしたい
}

func (u UserTable) Validate() error {
	if !strings.Contains(u.Email, "@") { // とりあえず
		return errors.New(fmt.Sprintf("Unexpected format of email: %v", u.Email))
	}
	return nil
}
