package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Code            string      `json:"code" form:"code" gorm:"type:varchar(50)"`
	Name            string      `json:"name" form:"name" gorm:"type:varchar(50)"`
	Description     string      `json:"description" form:"description" gorm:"type:text"`
	Product_Type_ID uint        `json:"product_type_id" form:"product_type_id"`
	Stock           uint        `json:"stock" form:"stock"`
	Price           uint        `json:"price" form:"price" gorm:"type:double"`
	ProductType     ProductType `gorm:"foreignKey:Product_Type_ID"`
}

type ProductType struct {
	gorm.Model
	Name string `json:"name" form:"name" gorm:"type:varchar(20)"`
}

type ProductResponse struct {
	ID              uint
	Code            string
	Name            string
	Description     string
	Product_Type_ID uint
	Stock           uint
	Price           uint
}

type AllProductResponse struct {
	ID    uint
	Name  string
	Price uint
}
