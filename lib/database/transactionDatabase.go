package database

import (
	"mini-project-apotek/config"
	"mini-project-apotek/models"
)

func SaveShipping(shipping *models.Shipping) error {
	if err := config.DB.Save(shipping).Error; err != nil {
		return err
	}
	return nil

}

func SaveTransaction(transaction *models.Transaction) error {
	if err := config.DB.Save(transaction).Error; err != nil {
		return err
	}
	return nil
}

func SaveTransactionDetail(transactionDetail *models.TransactionDetail) error {
	if err := config.DB.Save(transactionDetail).Error; err != nil {
		return err
	}
	return nil
}
