package handler

import (
	"log"
	"net/http"

	"github.com/komugi8/clava/domain"
	"github.com/komugi8/clava/usecase"
	"github.com/labstack/echo/v4"
)

type UserRegisterHandler struct {
	u *usecase.UserRegister
}

type UserRegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func NewUserRegisterHandler(u *usecase.UserRegister) *UserRegisterHandler {
	return &UserRegisterHandler{u: u}
}

func (h *UserRegisterHandler) RegisterUser(c echo.Context) error {
	req := &UserRegisterRequest{}
	if err := c.Bind(req); err != nil {
		return err
	}
	user := domain.User{
		Name:     req.Name,
		Password: req.Password,
		Role:     req.Role,
	}
	_, err := h.u.RegisterUser(c.Request().Context(), &user)
	if err != nil {
		log.Println("register user error: ", err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}
