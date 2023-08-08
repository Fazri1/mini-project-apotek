package database

import (
	"mini-project-apotek/config"
	"mini-project-apotek/models"
)

func SaveProductType(productType *models.ProductType) error {
	if err := config.DB.Save(productType).Error; err != nil {
		return err
	}

	return nil
}

func GetAllProductTypes() ([]models.ProductType, error) {
	var productTypes []models.ProductType
	if err := config.DB.Find(&productTypes).Error; err != nil {
		return nil, err
	}

	return productTypes, nil
}

func GetProductTypeByID(id string) (models.ProductType, error) {
	var productType models.ProductType
	if err := config.DB.First(&productType, id).Error; err != nil {
		return productType, err
	}

	return productType, nil
}

func DeleteProductType(id string) error {
	var productType models.ProductType
	if err := config.DB.Delete(&productType, id).Error; err != nil {
		return err
	}

	return nil
}