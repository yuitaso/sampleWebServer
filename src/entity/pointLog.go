package entity

import (
	"time"
)

type PointLog struct {
	Id         uint `gorm:"primarykey"`
	UserId     uint `gorm:"index:idx_user"`
	Amount     int
	ContractId uint
	CreatedAt  time.Time
}
