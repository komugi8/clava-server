package utils

import (
	"fmt"

	"github.com/komugi8/clava/domain"
)

type UserAlreadyExists struct {
	ID domain.UserID
}

func (e *UserAlreadyExists) Error() string {
	return fmt.Sprintf("User %d already exists", e.ID)
}

func ErrUserAlreadyExists(id domain.UserID) error {
	return &UserAlreadyExists{ID: id}
}
