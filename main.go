package main

import (
	"context"
	"net/http"

	"github.com/komugi8/clava/api/route"
	"github.com/komugi8/clava/pkg/config"
	"github.com/komugi8/clava/pkg/database"
	"github.com/labstack/echo/v4"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	db, err := database.NewDB(ctx, cfg)
	if err != nil {
		panic(err)
	}
	e := echo.New()
	user := e.Group("/user")
	route.NewSignUpRoute(db, user)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
