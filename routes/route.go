package routes

import (
	"mini-project-apotek/constants"
	"mini-project-apotek/controllers"
	mid "mini-project-apotek/middlewares"

	jwtMid "github.com/labstack/echo-jwt"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)
	e.GET("/product-types", controllers.GetProductTypesController)

	eJWT := e.Group("/jwt")
	eJWT.Use(jwtMid.JWT([]byte(constants.SECRET_KEY)))
	eJWT.POST("/products", controllers.AddProductController)
	eJWT.POST("/product-types", controllers.AddProductTypeController)
	eJWT.PUT("/product-types/:id", controllers.UpdateProductTypesController)
	eJWT.DELETE("/product-types/:id", controllers.DeleteProductTypeController)
	return e
}
