package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionNumber string `json:"transaction_number" form:"transaction_number" gorm:"unique"`
	Date              string `json:"date" form:"date" gorm:"type:timestamp"`
	UserID            uint   `json:"user_id" form:"user_id"`
	TotalQTY          uint   `json:"total_qty" form:"total_qty"`
	ShippingCost      uint   `json:"shipping_cost" form:"shipping_cost" gorm:"type:double"`
	TotalPrice        uint   `json:"total_price" form:"total_price" gorm:"type:double"`
	Status            string `json:"status" form:"status"`
	User              User   `gorm:"foreignKey:UserID"`
}

type TransactionDetail struct {
	ID            uint        `gorm:"primarykey"`
	TransactionID uint        `json:"transaction_id"`
	ProductID     uint        `json:"product_id"`
	QTY           uint        `json:"qty"`
	Price         uint        `json:"price"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	Product       Product     `gorm:"foreignKey:ProductID"`
}

type Shipping struct {
	ID          uint   `gorm:"primarykey;autoIncrement"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}

type CheckOut struct {
	Address struct {
		Detail     string `json:"detail" form:"detail"`
		Province   string `json:"province" form:"province"`
		City       string `json:"city" form:"city"`
		District   string `json:"district" form:"district"`
		PostalCode string `json:"postal_code" form:"postal_code"`
	} `json:"address" form:"address"`
	Name        string `json:"name" form:"name"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
	// Products    []struct {
	ProductID uint `json:"product_id" form:"product_id"`
	QTY       uint `json:"qty" form:"qtr"`
	// } `json:"products"`

	PaymentMethod string `json:"payment_method" form:"payment_method"`
}

type CheckOutResponse struct {
	SubtotalsForProduct  uint `json:"subtotals_for_product"`
	SubtotalsForShipping uint `json:"subtotals_for_shipping"`
	TotalPayment         uint `json:"total_payment"`
}
