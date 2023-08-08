package middlewares

import (
	"mini-project-apotek/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint, name, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.JWT_SECRET_KEY))
}

func CheckTokenRole(tokenString string) (bool, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JWT_SECRET_KEY), nil
	})
	if err != nil {
		return false, err
	}

	if claims["role"] == "admin" {
		return true, nil
	}
	return false, nil
}