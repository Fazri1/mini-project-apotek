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

func SaveOrder(order *models.Order) error {
	if err := config.DB.Save(order).Error; err != nil {
		return err
	}

	return nil
}

func SaveOrderDetail(orderDetail *models.OrderDetail) error {
	if err := config.DB.Save(orderDetail).Error; err != nil {
		return err
	}

	return nil
}

func GetUserOrders(user_id string) ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Preload("User").Joins("User").Where("user_id = ?", user_id).Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func GetUserOrderDetail(user_id, order_id string) (models.OrderDetail, error) {
	var orderDetail models.OrderDetail
	if err := config.DB.Preload("Order").Preload("Product").Preload("Shipping").Where("order_id = ?", order_id).Find(&orderDetail).Error; err != nil {
		return orderDetail, err
	}

	return orderDetail, nil
}

func UpdateOrderPayment(orderUpdate *models.Notification) error {
	var order models.Order
	if err := config.DB.Model(&order).Where("order_number = ?", orderUpdate.OrderID).Updates(models.Order{Status: "packed",
		PaymentStatus: orderUpdate.TransactionStatus, PaymentMethod: orderUpdate.PaymentType}).Error; err != nil {
		return err
	}

	return nil
}

func GetAllOrders() ([]models.Order, error) {
	var orders []models.Order
	if err := config.DB.Preload("User").Find(&orders).Error; err != nil {
		return orders, err
	}

	return orders, nil
}

func GetOrderDetail(order_id string) (models.OrderDetail, error) {
	var orderDetail models.OrderDetail
	if err := config.DB.Preload("Order").Preload("Product").Preload("Shipping").Where("order_id = ?", order_id).First(&orderDetail).Error; err != nil {
		return orderDetail, err
	}

	return orderDetail, nil
}

func UpdateStatusOrder(status, id string) error {
	var order models.Order
	if err := config.DB.Model(&order).Where("id = ?", id).Update("status", status).Error; err != nil {
		return err
	}

	return nil
}