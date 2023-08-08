package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderNumber   string `json:"order_number" form:"order_number" gorm:"unique"`
	Date          string `json:"date" form:"date" gorm:"type:datetime(3)"`
	UserID        uint   `json:"user_id" form:"user_id"`
	TotalQTY      uint   `json:"total_qty" form:"total_qty"`
	ShippingCost  uint   `json:"shipping_cost" form:"shipping_cost" gorm:"type:double"`
	TotalPrice    uint   `json:"total_price" form:"total_price" gorm:"type:double"`
	Status        string `json:"status" form:"status"`
	PaymentMethod string `json:"payment_method" form:"payment_method"`
	PaymentStatus string `json:"payment_status" form:"payment_status"`
	SnapToken     string `json:"snap_token"`
	User          User   `gorm:"foreignKey:UserID"`
}

type OrderResponse struct {
	ID            uint
	OrderNumber   string
	Date          string
	UserID        uint
	TotalQTY      uint
	ShippingCost  uint
	TotalPrice    uint
	Status        string
	PaymentMethod string
	PaymentStatus string
	User          UserResponse
}

type MidtransRequest struct {
	OrderNumber  string
	Amount       int64
	Product      AllProductResponse
	QTY          int32
	ShippingCost int64
	User         struct {
		Name  string
		Email string
		Phone string
	}
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
	ProductID   uint   `json:"product_id" form:"product_id"`
	QTY         uint   `json:"qty" form:"qty"`
}

type CheckOutResponse struct {
	SubtotalsForProduct  uint `json:"subtotals_for_product"`
	SubtotalsForShipping uint `json:"subtotals_for_shipping"`
	TotalPayment         uint `json:"total_payment"`
}