package usecase

import (
	"context"

	"github.com/komugi8/clava/domain"
	// "github.com/komugi8/clava/utils"
)

type UserRegister struct {
	userRepo domain.UserRepository
}

func NewUserRegister(userRepo domain.UserRepository) *UserRegister {
	return &UserRegister{userRepo: userRepo}
}

func (u *UserRegister) RegisterUser(c context.Context, user *domain.User) (*domain.User, error) {
	// if _, err := u.userRepo.GetUserByID(c, user.ID); err == nil {
	// 	return nil, utils.ErrUserAlreadyExists(user.ID)
	// }
	createdUser, err := u.userRepo.CreateUser(c, user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
