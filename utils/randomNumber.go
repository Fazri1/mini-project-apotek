package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomTransactionID() string {
	currentTime := time.Now()
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9999) + 1000
	transactionID := currentTime.Format("02012006") + "-" + fmt.Sprintf("%d", randomNum)

	return transactionID
}