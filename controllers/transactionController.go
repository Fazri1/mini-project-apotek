package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/lib/services"
	"mini-project-apotek/middlewares"
	"mini-project-apotek/models"
	"mini-project-apotek/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func CheckOutController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if !admin {
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

		userID, _ := strconv.Atoi(c.Param("userID"))
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
		transactionDetail.ShippingID = shipping.ID
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
	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized Action",
	})
}

func GetTransactionsController(c echo.Context) error {
	transactions, err := database.GetUserTransactions(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}
	transactionsList := []models.TransactionResponse{}
	for i := range transactions {
		transactionsResponse := models.TransactionResponse{}
		transactionsResponse.ID = transactions[i].ID
		transactionsResponse.TransactionNumber = transactions[i].TransactionNumber
		transactionsResponse.Date = transactions[i].Date
		transactionsResponse.UserID = transactions[i].UserID
		transactionsResponse.TotalQTY = transactions[i].TotalQTY
		transactionsResponse.ShippingCost = transactions[i].ShippingCost
		transactionsResponse.TotalPrice = transactions[i].TotalPrice
		transactionsResponse.Status = transactions[i].Status
		transactionsResponse.User.ID = transactions[i].User.ID
		transactionsResponse.User.Name = transactions[i].User.Name
		transactionsResponse.User.Email = transactions[i].User.Email
		transactionsList = append(transactionsList, transactionsResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Success Get Transactions",
		"transactions": transactionsList,
	})
}

func GetUserTransactionDetailController(c echo.Context) error {
	transactionDetail, err := database.GetUserTransactionDetail(c.Param("userID"), c.Param("transactionID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	transactionDetailResponse := models.TransactionDetailResponse{ID: transactionDetail.ID, TransactionID: transactionDetail.TransactionID,
		ProductID: transactionDetail.ProductID, QTY: transactionDetail.QTY, Price: transactionDetail.Price,
		Transaction: struct {
			TransactionNumber string
			Date              string
		}{transactionDetail.Transaction.TransactionNumber, transactionDetail.Transaction.Date},
		Product: models.AllProductResponse{transactionDetail.Product.ID, transactionDetail.Product.Name, transactionDetail.Product.Price},
		Address: transactionDetail.Shipping}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Success Get Transaction Detail",
		"transactions": transactionDetailResponse,
	})
}
