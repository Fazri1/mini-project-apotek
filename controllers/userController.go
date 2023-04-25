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
	err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user":    user,
	})
}

func LoginController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := database.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "failed login",
		})
	}

	token, err := middlewares.CreateToken(user.ID, user.Name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userResponse := models.UserResponse{user.ID, user.Name, user.Email, token}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    userResponse,
	})
}
