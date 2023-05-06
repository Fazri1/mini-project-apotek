package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NotificationController(c echo.Context) error {
	var notification models.Notification
	c.Bind(&notification)

	err := database.UpdateTransactionPayment(&notification)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success",
	})
}
