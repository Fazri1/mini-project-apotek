package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomOrderID() string {
	currentTime := time.Now()
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9999) + 1000
	orderID := currentTime.Format("02012006") + "-" + fmt.Sprintf("%d", randomNum)

	return orderID
}

func GenerateRandomString(name string) string {
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9999) + 1000
	imageName := name + "-" + fmt.Sprintf("%d", randomNum)

	return imageName
}
