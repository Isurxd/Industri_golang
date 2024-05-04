package routes

import (
	"crud-api/controllers"

	"github.com/labstack/echo/v4"
)

func Routes() (e *echo.Echo) {
	e = echo.New()

	productRoutes := e.Group("/product")
	productRoutes.GET("", controllers.ReadAllProducts)
	productRoutes.POST("/create", controllers.CreateProduct)
	productRoutes.GET("/:id", controllers.ReadDetailProducts)
	productRoutes.PUT("/update", controllers.UpdateProduct)
	productRoutes.DELETE("/:id", controllers.DeleteProduct)

	categoryRoutes := e.Group("/category")
	categoryRoutes.GET("", controllers.ReadAllCategorys)
	categoryRoutes.POST("/create", controllers.CreateCategorys)
	categoryRoutes.GET("/:id", controllers.ReadDetailCategorys)
	categoryRoutes.PUT("/update", controllers.UpdateCategorys)
	categoryRoutes.DELETE("/:id", controllers.DeleteCategorys)
	return
}
