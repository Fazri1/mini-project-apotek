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

type TransactionDetail struct {
	ID            uint        `gorm:"primarykey"`
	TransactionID uint        `json:"transaction_id"`
	ProductID     uint        `json:"product_id"`
	QTY           uint        `json:"qty"`
	Price         uint        `json:"price"`
	ShippingID    uint        `json:"shipping_id" form:"shipping_id"`
	Transaction   Transaction `gorm:"foreignKey:TransactionID"`
	Product       Product     `gorm:"foreignKey:ProductID"`
	Shipping      Shipping    `gorm:"foreignKey:ShippingID"`
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

type TransactionDetailResponse struct {
	ID            uint
	TransactionID uint
	ProductID     uint
	QTY           uint
	Price         uint
	Transaction   struct {
		TransactionNumber string
		Date              string
		Status            string
	}
	Product AllProductResponse
	Address Shipping
}

type MidtransRequest struct {
	TransactionNumber string
	Amount            int64
	Product           AllProductResponse
	QTY               int32
	ShippingCost      int64
	User              struct {
		Name  string
		Email string
		Phone string
	}
}