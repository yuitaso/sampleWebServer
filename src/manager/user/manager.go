package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/env"
	"github.com/yuitaso/sampleWebServer/src/util"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type UserTable struct {
	gorm.Model
	IdHash   string
	Email    string
	Password string
}

func (u UserTable) TableName() string {
	return "user"
}

func Create(email string, rawPass string) (uint, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.MinCost)
	idHash, err := util.GenerateRandomHash()

	fmt.Println(idHash)

	if err != nil {
		return 0, err
	}

	data := UserTable{IdHash: idHash, Email: email, Password: string(passwordHash)}
	if err = data.Validate(); err != nil {
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

	return entity.User{IdHash: result.IdHash, Email: result.Email}, nil
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
