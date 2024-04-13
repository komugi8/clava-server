package domain

import (
	"context"
)

type UserRepository interface {
	GetUsersByIDs(c context.Context, userIDs []UserID) (*[]User, error)
	GetUserByID(c context.Context, userID UserID) (*User, error)
	CreateUser(c context.Context, user *User) (*User, error)
	DeleteUser(c context.Context, userID UserID) (*User, error)
	UpdateUser(c context.Context, user *User) (*User, error)
}
