package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	TransactionNumber string `json:"transaction_number" form:"transaction_number" gorm:"unique"`
	Date              string `json:"date" form:"date" gorm:"type:datetime(3)"`
	UserID            uint   `json:"user_id" form:"user_id"`
	TotalQTY          uint   `json:"total_qty" form:"total_qty"`
	ShippingCost      uint   `json:"shipping_cost" form:"shipping_cost" gorm:"type:double"`
	TotalPrice        uint   `json:"total_price" form:"total_price" gorm:"type:double"`
	PaymentMethod     string `json:"payment_method" form:"payment_method"`
	PaymentStatus     string `json:"payment_status" form:"payment_status"`
	SnapToken         string `json:"snap_token"`
	User              User   `gorm:"foreignKey:UserID"`
}

type TransactionResponse struct {
	ID                uint
	TransactionNumber string
	Date              string
	UserID            uint
	TotalQTY          uint
	ShippingCost      uint
	TotalPrice        uint
	PaymentMethod     string
	PaymentStatus     string
	User              UserResponse
}