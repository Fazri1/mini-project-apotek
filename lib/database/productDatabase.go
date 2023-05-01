package database

import (
	"mini-project-apotek/config"
	"mini-project-apotek/models"
)

func AddProduct(product *models.Product) error {
	if err := config.DB.Save(product).Error; err != nil {
		return err
	}
	return nil
}
