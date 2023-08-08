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

	e.POST("/register", controllers.RegisterController)
	e.POST("/login", controllers.LoginController)
	e.GET("/product-types", controllers.GetProductTypesController)
	e.GET("/products", controllers.GetProductsController)
	e.GET("/products/:id", controllers.GetProductDetailController)
	e.GET("/products/search", controllers.SearchProductController)
	e.POST("/notification", controllers.NotificationController)

	eJWT := e.Group("/auth")
	eJWT.Use(jwtMid.JWT([]byte(constants.JWT_SECRET_KEY)))
	eJWT.PUT("/users/:id", controllers.UpdateUserController)
	eJWT.POST("/product-types", controllers.AddProductTypeController)
	eJWT.PUT("/product-types/:id", controllers.UpdateProductTypeController)
	eJWT.DELETE("/product-types/:id", controllers.DeleteProductTypeController)
	eJWT.POST("/products", controllers.AddProductController)
	eJWT.PUT("/products/:id", controllers.UpdateProductController)
	eJWT.DELETE("/products/:id", controllers.DeleteProductController)
	eJWT.POST("/users/:userID/checkout", controllers.CheckOutController)
	eJWT.GET("/users/:userID/orders", controllers.GetUserOrdersController)
	eJWT.GET("/users/:userID/orders/:orderID", controllers.GetUserOrderDetailController)
	eJWT.PUT("/users/:userID/orders/:orderID", controllers.UpdateStatusOrderController)
	eJWT.GET("/orders", controllers.GetAllOrdersController)
	eJWT.GET("/orders/:orderID", controllers.GetOrderDetailController)
	eJWT.PUT("/orders/:orderID", controllers.UpdateStatusOrderController)

	return e
}