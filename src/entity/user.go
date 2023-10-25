package entity

import "github.com/google/uuid"

type User struct {
	Id uint
	Uuid  uuid.UUID
	Email string
}

var CtxAuthUserKey = "authUser"
var InitialPointAmount = 10000 // when user created
