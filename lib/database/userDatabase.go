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

func Login(user *models.User) (string, error) {
	var userForLogin models.User
	if err := config.DB.Table("users").Select("password").Where("email = ?", user.Email).First(&userForLogin).Error; err != nil {
		return "", err
	}

	err := utils.ComparePassword(userForLogin.Password, user.Password)
	if err != nil {
		return "", err
	}

	if err := config.DB.Where("email = ? AND password = ?", user.Email, userForLogin.Password).First(&user).Error; err != nil {
		return "", err
	}

	return user.Role, nil
}

func GetUserEmail(id string) string {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return ""
	}

	return user.Email
}

func UpdateUser(id string, updatedUser *models.User) error {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return err
	}

	user.Name = updatedUser.Name
	user.Email = updatedUser.Email
	user.Password = updatedUser.Password

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = utils.ComparePassword(hashedPassword, user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}
