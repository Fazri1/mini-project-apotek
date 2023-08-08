package models

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

type TransactionDetailResponse struct {
	ID            uint
	TransactionID uint
	ProductID     uint
	QTY           uint
	Price         uint
	Transaction   struct {
		TransactionNumber string
		Date              string
		PaymentStatus     string
	}
	Product AllProductResponse
	Address Shipping
}