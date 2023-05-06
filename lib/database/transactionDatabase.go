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

func GetUserTransactions(user_id string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := config.DB.Preload("User").Joins("User").Where("user_id = ?", user_id).Find(&transactions).Error; err != nil {
		return transactions, err
	}
	return transactions, nil
}

func GetUserTransactionDetail(user_id, transaction_id string) (models.TransactionDetail, error) {
	var transactionDetail models.TransactionDetail
	if err := config.DB.Preload("Transaction").Preload("Product").Preload("Shipping").Where("transaction_id = ?", transaction_id).Find(&transactionDetail).Error; err != nil {
		return transactionDetail, err
	}
	return transactionDetail, nil
}

func UpdateTransactionPayment(transactionUpdate *models.Notification) error {
	var transaction models.Transaction
	if err := config.DB.Model(&transaction).Where("transaction_number = ?", transactionUpdate.OrderID).Updates(models.Transaction{PaymentStatus: transactionUpdate.TransactionStatus, PaymentMethod: transactionUpdate.PaymentType}).Error; err != nil {
		return err
	}
	return nil
}
