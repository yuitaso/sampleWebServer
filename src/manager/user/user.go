package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/yuitaso/sampleWebServer/src/entity"
	"github.com/yuitaso/sampleWebServer/src/manager"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	pointLogManager "github.com/yuitaso/sampleWebServer/src/manager/pointLog"
)

type UserTable struct {
	gorm.Model
	Uuid     string
	Email    string
	Password string
}

func (u UserTable) TableName() string {
	return "user"
}

func Insert(email string, rawPass string) (*entity.User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(rawPass), bcrypt.MinCost)
	user_uuid, err := uuid.NewRandom()
	if err != nil {
		return &entity.User{}, err
	}

	data := UserTable{Uuid: user_uuid.String(), Email: email, Password: string(passwordHash)}
	// # if validate model
	// if err = data.Validate(); err != nil {
	// 	return 0, err
	// }

	if executed := manager.DB.Create(&data); executed.Error != nil {
		return &entity.User{}, executed.Error
	}

	return &entity.User{Id: data.ID, Uuid: uuid.MustParse(data.Uuid), Email: data.Email}, nil
}

func FindById(id int) (entity.User, error) {
	var result UserTable
	if executed := manager.DB.First(&result, id); executed.Error != nil {
		return entity.User{}, executed.Error
	}

	return entity.User{Uuid: uuid.MustParse(result.Uuid), Email: result.Email}, nil
}

func VerifyAndGetUser(email string, password string) (*entity.User, error) {
	var result UserTable
	if executed := manager.DB.First(&result, "email = ?", email); executed.Error != nil {
		return &entity.User{}, executed.Error
	}

	// imo 毎回キャストするとメモリ効率が悪いのでいい感じにしたい
	if err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return &entity.User{}, err
	}

	return &entity.User{Id: result.ID, Email: result.Email, Uuid: uuid.MustParse(result.Uuid)}, nil
}

func (u UserTable) Validate() error {
	if !strings.Contains(u.Email, "@") { // 条件は適当
		return errors.New(fmt.Sprintf("Unexpected format of email: %v", u.Email))
	}
	return nil
}

func CreateUserWithPointGrant(email string, rawPass string, amount int) (*entity.User, error) {
	var newUser *entity.User
	var err error
	err = manager.DB.Transaction(func(db *gorm.DB) error {
		var err error
		newUser, err = Insert(email, rawPass)
		if err != nil {
			return err
		}

		if err := pointLogManager.Insert(newUser, amount); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
