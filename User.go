package model

type User struct {
	Username string
}

func NewUser(username string) *User {
	return &User{Username: username}
}
