package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/middlewares"
	"mini-project-apotek/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddProductTypeController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		var productType models.ProductType
		c.Bind(&productType)

		err = database.SaveProductType(&productType)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":      "Success Add Product Type",
			"product type": productType,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})
}

func GetProductTypesController(c echo.Context) error {
	productTypes, err := database.GetAllProductTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Success Get All Product Types",
		"product types": productTypes,
	})

}

func UpdateProductTypeController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		var updatedProductType models.ProductType
		c.Bind(&updatedProductType)

		productType, err := database.GetProductTypeByID(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err.Error(),
			})
		}

		productType.Name = updatedProductType.Name
		err = database.SaveProductType(&productType)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":      "success update product type",
			"product type": productType,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})

}

func DeleteProductTypeController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		err := database.DeleteProductType(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Success Delete Product Type",
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})
}
