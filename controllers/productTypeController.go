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

		productTypeResponse := models.ProductTypeResponse{ID: productType.ID, Name: productType.Name}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":      "Success add product type",
			"product type": productTypeResponse,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}

func GetProductTypesController(c echo.Context) error {
	productTypes, err := database.GetAllProductTypes()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	productTypesResponseList := []models.ProductTypeResponse{}

	for i := range productTypes {
		var productTypeResponse models.ProductTypeResponse
		productTypeResponse.ID = productTypes[i].ID
		productTypeResponse.Name = productTypes[i].Name
		productTypesResponseList = append(productTypesResponseList, productTypeResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":       "Success get all product types",
		"product types": productTypesResponseList,
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
			"message": "Success update product type",
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
			"message": "Success delete product type",
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}
