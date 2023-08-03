package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
}
