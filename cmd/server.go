package cmd

import (
	"crud_product/config"
	"crud_product/pkg/router"
	database "crud_product/shared/db"
	"github.com/labstack/echo/v4"
)

func RunServe() {
	e := echo.New()
	g := e.Group("")
	conf := config.GetConfig()

	Apply(e, g, conf)
	e.Logger.Fatal(e.Start(":5500"))
}

func Apply(e *echo.Echo, g *echo.Group, conf config.Configuration) {
	db := database.NewInstanceDB(conf)
	router.NewProductRouter(e, g, db)
}
