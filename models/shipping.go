package models

type Shipping struct {
	ID          uint   `gorm:"primarykey;autoIncrement"`
	Name        string `json:"name" form:"name"`
	Address     string `json:"address" form:"address"`
	PhoneNumber string `json:"phone_number" form:"phone_number"`
}