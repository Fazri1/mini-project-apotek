package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomTransactionID() string {
	// Mengambil waktu saat ini sebagai string
	currentTime := time.Now()

	// Menghasilkan bilangan acak antara 1000 dan 9999
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Intn(9999) + 1000

	// Menggabungkan waktu dan bilangan acak menjadi nomor transaksi acak
	transactionID := currentTime.Format("02012006") + "-" + fmt.Sprintf("%d", randomNum)

	return transactionID
}
