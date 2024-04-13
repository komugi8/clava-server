package route

import (
	"github.com/jmoiron/sqlx"
	"github.com/komugi8/clava/api/handler"
	"github.com/komugi8/clava/infra"
	"github.com/komugi8/clava/usecase"
	"github.com/labstack/echo/v4"
)

func NewSignUpRoute(db *sqlx.DB, group *echo.Group) {
	ur := infra.NewUserRepository(db)
	sh := handler.NewUserRegisterHandler(usecase.NewUserRegister(ur))
	group.POST("/signup", func(c echo.Context) error {
		return sh.RegisterUser(c)
	})
}