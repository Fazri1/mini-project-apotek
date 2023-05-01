package database

import (
	"mini-project-apotek/config"
	"mini-project-apotek/models"
)

func SaveProduct(product *models.Product) error {
	if err := config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}

func GetAllProducts() ([]models.AllProductResponse, error) {
	var products []models.AllProductResponse

	if err := config.DB.Table("products").Select("id, name").Scan(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}

func GetProductById(id string) (models.Product, error) {
	var product models.Product
	if err := config.DB.Preload("ProductType").First(&product, id).Error; err != nil {
		return product, err
	}

	return product, nil
}

func DeleteProduct(id string) error {
	var product models.Product
	if err := config.DB.Delete(&product, id).Error; err != nil {
		return err
	}
	return nil
}
