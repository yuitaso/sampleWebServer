package entity

import (
	"time"
)

type PointLog struct {
	ID         uint `gorm:"primarykey"`
	UserId     uint `gorm:"index:idx_user"`
	Amount     int
	ContractId uint
	CreatedAt  time.Time
}
