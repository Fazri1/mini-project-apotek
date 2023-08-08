package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name" gorm:"type:varchar(255)"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role" gorm:"type:varchar(10)"`
}

type UserResponse struct {
	ID    uint
	Name  string
	Email string
}