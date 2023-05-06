package models

type Notification struct {
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	TransactionStatus string `json:"transaction_status"`
}
