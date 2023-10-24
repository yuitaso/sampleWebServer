package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Id     uint
	Uuid   uuid.UUID
	Price  int
	UserId uint
}

func (item *Item) IsOwner(user *User) bool {
	return item.UserId == user.Id
}
