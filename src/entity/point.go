package entity

import "gorm.io/gorm"

type UserPoint struct {
	gorm.Model
	Amount int
	UserID uint
}

type PointLog struct {
	gorm.Model
	UserId uint
	Amount int
	ContractId uint
}
