package router

import (
	"crud_product/pkg/controller"
	"crud_product/pkg/repository"
	"crud_product/pkg/usecase"
	"database/sql"
	"github.com/labstack/echo/v4"
)

func NewProductRouter(e *echo.Echo, g *echo.Group, db *sql.DB) {
	pr := repository.NewUserRepository(db)
	pu := usecase.NewProductUsecase(pr)
	pc := controller.NewProductController(pu)

	e.GET("/product", pc.GetAllProduct)
	e.POST("/product", pc.CreateProduct)
	e.PUT("/product/:id", pc.UpdateProduct)
	e.DELETE("/product/:id", pc.DeleteProduct)
	e.GET("/product/:id", pc.GetProductById)
}
