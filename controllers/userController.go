package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/middlewares"
	"mini-project-apotek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	user.Role = "customer"
	id := database.CheckEmail(user.Email)
	if id != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Email already registered",
		})
	}

	err := database.CreateUser(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	userResponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user":    userResponse,
	})
}

func LoginController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	role, err := database.Login(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Failed login",
		})
	}

	userResponse := models.UserResponse{ID: user.ID, Name: user.Name, Email: user.Email}

	token, err := middlewares.CreateToken(user.ID, user.Name, role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success login",
		"user":    userResponse,
		"token":   token,
	})
}

func UpdateUserController(c echo.Context) error {
	var user models.User
	c.Bind(&user)

	err := database.UpdateUser(c.Param("id"), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success update user",
	})
}