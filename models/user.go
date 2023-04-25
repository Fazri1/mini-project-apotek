package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email" gorm:"unique"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}

type UserResponse struct {
	ID    uint
	Name  string
	Email string
	Token string
}
