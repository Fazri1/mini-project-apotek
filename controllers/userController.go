package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/middlewares"
	"mini-project-apotek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	var user models.User

	c.Bind(&user)
	user.Role = "customer"

	err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	user.Password = "$"
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user":    user,
	})
}

func LoginController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	role, err := database.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed login",
		})
	}

	userResponse := models.UserResponse{user.ID, user.Name, user.Email}
	// if user.Role == "admin" {
	token, err := middlewares.CreateToken(user.ID, user.Name, role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
		"token":   token,
	})
	// }
	// return c.JSON(http.StatusOK, map[string]interface{}{
	// 	"message": "success login",
	// 	"user":    userResponse,
	// })
}
