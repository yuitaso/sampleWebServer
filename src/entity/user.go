package entity

type User struct {
	Id uint
	Uuid  string
	Email string
}

var CtxAuthUserKey = "authUser"
