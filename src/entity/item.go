package entity

import (
	"github.com/google/uuid"
)

type Item struct {
	Id           uint
	Uuid         uuid.UUID
	Price        int
	SellerUserId uint
}

func (item *Item) IsOwner(user *User) bool {
	return item.SellerUserId == user.Id
}
