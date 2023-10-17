package user

func Create() (User, error) {
	u := User{Name: "manager created", Password: "pass"}
	return u, nil
}
