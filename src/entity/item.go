package entity

import "github.com/google/uuid"

type Item struct {
	Id uint
	Uuid uuid.UUID
	Price int
	UserId uint	
}