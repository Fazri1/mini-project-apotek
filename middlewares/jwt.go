package middlewares

import (
	"mini-project-apotek/constants"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(id uint, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_KEY))
}
