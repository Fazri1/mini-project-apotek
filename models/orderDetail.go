package models

type OrderDetail struct {
	ID         uint     `gorm:"primarykey"`
	OrderID    uint     `json:"order_id"`
	ProductID  uint     `json:"product_id"`
	QTY        uint     `json:"qty"`
	Price      uint     `json:"price"`
	ShippingID uint     `json:"shipping_id" form:"shipping_id"`
	Order      Order    `gorm:"foreignKey:OrderID"`
	Product    Product  `gorm:"foreignKey:ProductID"`
	Shipping   Shipping `gorm:"foreignKey:ShippingID"`
}

type OrderDetailResponse struct {
	ID        uint
	OrderID   uint
	ProductID uint
	QTY       uint
	Price     uint
	Order     struct {
		OrderNumber   string
		Date          string
		Status        string
		PaymentStatus string
	}
	Product AllProductResponse
	Address Shipping
}