package controllers

import (
	"mini-project-apotek/lib/database"
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
