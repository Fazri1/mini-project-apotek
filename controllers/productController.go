package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddProductController(c echo.Context) error {
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

func DeleteProductController(c echo.Context) error {
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
