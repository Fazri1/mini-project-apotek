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
