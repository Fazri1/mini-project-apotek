package controllers

import (
	"mini-project-apotek/lib/database"
	awss3 "mini-project-apotek/lib/services/aws"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		var product models.Product
		var imageURI string
		var err error
		c.Bind(&product)

		image, err := c.FormFile("image")
		if image != nil {
			// os.Setenv("AWS_ACCESS_KEY_ID", constants.AWS_ACCESS_KEY_ID)
			// os.Setenv("AWS_SECRET_ACCESS_KEY", constants.AWS_SECRET_ACCESS_KEY)

			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]string{
					"message": err.Error(),
				})
			}

			imageURI, err = awss3.UploadFileS3(product.Name, image)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": err.Error(),
				})
			}
		}
		product.ImageURI = imageURI
		err = database.SaveProduct(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		productResponse := models.ProductResponse{ID: product.ID, Code: product.Code, Name: product.Name, Description: product.Description,
			ProductTypeID: product.ProductTypeID, Stock: product.Stock, Price: product.Price, ImageURI: product.ImageURI}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success add product",
			"product": productResponse,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})

}

func GetProductsController(c echo.Context) error {
	products, err := database.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "Success get all products",
		"products": products,
	})

}

func GetProductDetailController(c echo.Context) error {
	product, err := database.GetDetailProduct(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Failed get product detail",
		})
	}

	productDetailResponse := models.ProductDetailResponse{ID: product.ID, Code: product.Code,
		Name: product.Name, Description: product.Description, ProductTypeID: product.ProductTypeID, Stock: product.Stock,
		Price: product.Price, ImageURI: product.ImageURI, ProductType: models.ProductTypeResponse{ID: product.ProductType.ID, Name: product.ProductType.Name}}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get product detail",
		"product": productDetailResponse,
	})
}

func UpdateProductController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		// os.Setenv("AWS_ACCESS_KEY_ID", constants.AWS_ACCESS_KEY_ID)
		// os.Setenv("AWS_SECRET_ACCESS_KEY", constants.AWS_SECRET_ACCESS_KEY)
		var updatedProduct models.Product
		c.Bind(&updatedProduct)

		product, err := database.GetProductById(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		imageURI := product.ImageURI
		image, _ := c.FormFile("image")
		if image != nil {
			uri, s := awss3.UploadFileS3(updatedProduct.Name, image)
			if s != nil {
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"message": s.Error(),
				})
			}
			imageURI = uri
		}

		product.Code = updatedProduct.Code
		product.Name = updatedProduct.Name
		product.Description = updatedProduct.Description
		product.ProductTypeID = updatedProduct.ProductTypeID
		product.Stock = updatedProduct.Stock
		product.Price = updatedProduct.Price
		product.ImageURI = imageURI

		err = database.SaveProduct(&product)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": "Failed update product",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Success update product",
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}

func DeleteProductController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
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

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Success delete product",
		})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
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
		"message": "Success search product",
		"product": product,
	})
}