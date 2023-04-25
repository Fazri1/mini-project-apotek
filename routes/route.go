package routes

import (
	"mini-project-apotek/controllers"
	mid "mini-project-apotek/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	mid.LogMiddleware(e)

	e.POST("/register", controllers.CreateUserController)
	e.POST("/login", controllers.LoginController)

	return e
}
