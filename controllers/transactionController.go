package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/lib/services"
	"mini-project-apotek/models"
	"mini-project-apotek/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func CheckOutController(c echo.Context) error {
	var dataCheckOut models.CheckOut
	var shipping models.Shipping
	var transaction models.Transaction

	c.Bind(&dataCheckOut)

	cities := services.GetCityService()
	// provinces := services.GetProvinceService()
	city := strings.Title(dataCheckOut.Address.City)
	deliveryCost, err := services.GetDeliveryCostService(cities[city])
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	address := dataCheckOut.Address.Detail + ", " + dataCheckOut.Address.District + ", " + dataCheckOut.Address.City + ", " + dataCheckOut.Address.Province + ", " + dataCheckOut.Address.PostalCode
	shipping.Name = dataCheckOut.Name
	shipping.Address = address
	shipping.PhoneNumber = dataCheckOut.PhoneNumber

	err = database.SaveShipping(&shipping)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userID, _ := strconv.Atoi(c.Param("id"))
	product, _ := database.GetProductById(strconv.Itoa(int(dataCheckOut.ProductID)))
	totalProductPrice := product.Price * dataCheckOut.QTY
	totalPrice := totalProductPrice + uint(deliveryCost)

	num := utils.GenerateRandomTransactionID()
	transaction.TransactionNumber = num
	transaction.Date = time.Now().Format("2006-01-02 15:04:05")
	transaction.UserID = uint(userID)
	transaction.TotalQTY = dataCheckOut.QTY
	transaction.ShippingCost = uint(deliveryCost)
	transaction.TotalPrice = totalPrice
	transaction.Status = "Ok"
	errTransaction := database.SaveTransaction(&transaction)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": errTransaction.Error(),
		})
	}

	// for _, v := range dataCheckOut.Products {
	var transactionDetail models.TransactionDetail
	transactionDetail.TransactionID = transaction.ID
	transactionDetail.ProductID = dataCheckOut.ProductID
	transactionDetail.QTY = dataCheckOut.QTY
	product, _ = database.GetProductById(strconv.Itoa(int(dataCheckOut.ProductID)))
	transactionDetail.Price = product.Price
	errTD := database.SaveTransactionDetail(&transactionDetail)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": errTD.Error(),
		})
	}

	// }
	checkOutResponse := models.CheckOutResponse{totalProductPrice, uint(deliveryCost), totalPrice}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"Success": checkOutResponse,
	})
}
