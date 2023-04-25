package database

import (
	"mini-project-apotek/config"
	"mini-project-apotek/models"
	"mini-project-apotek/utils"
)

func CreateUser(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = utils.ComparePassword(hashedPassword, user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	if err := config.DB.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func Login(user *models.User) error {
	var userForLogin models.User
	if err := config.DB.Table("users").Select("password").Where("email = ?", user.Email).First(&userForLogin).Error; err != nil {
		return err
	}

	err := utils.ComparePassword(userForLogin.Password, user.Password)
	if err != nil {
		return err
	}

	if err := config.DB.Where("email = ? AND password = ?", user.Email, userForLogin.Password).First(&user).Error; err != nil {
		return err
	}
	return nil
}
