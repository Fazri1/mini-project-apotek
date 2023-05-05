package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/middlewares"
	"mini-project-apotek/models"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func AddProductController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		var product models.Product

		c.Bind(&product)
		err := database.SaveProduct(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}
		productResponse := models.ProductResponse{product.ID, product.Code, product.Name, product.Description, product.Product_Type_ID, product.Stock, product.Price}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Add Product",
			"product": productResponse,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})

}

func GetProductsController(c echo.Context) error {
	products, err := database.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success Get All Products",
		"products": products,
	})

}

func GetProductDetailController(c echo.Context) error {
	product, err := database.GetProductById(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed Get Product Detail",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Get Product Detail",
		"product": product,
	})
}

func UpdateProductController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		var updatedProduct models.Product
		c.Bind(&updatedProduct)

		product, err := database.GetProductById(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		product.Code = updatedProduct.Code
		product.Name = updatedProduct.Name
		product.Description = updatedProduct.Description
		product.Stock = updatedProduct.Stock
		product.Price = updatedProduct.Price

		err = database.SaveProduct(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Failed Update Product",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success Update Product",
			"product": product,
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})
}

func DeleteProductController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		err := database.DeleteProduct(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Success Delete Product",
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})
}

func SearchProductController(c echo.Context) error {
	product, err := database.SearchProduct(c.QueryParam("keyword"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success Search Product",
		"product": product,
	})
}
