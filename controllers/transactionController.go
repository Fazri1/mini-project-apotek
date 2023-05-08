package controllers

import (
	"mini-project-apotek/lib/database"
	"mini-project-apotek/lib/services/midtrans"
	servicesRO "mini-project-apotek/lib/services/rajaongkir"
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
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if !admin {
		var dataCheckOut models.CheckOut
		var shipping models.Shipping
		var transaction models.Transaction

		c.Bind(&dataCheckOut)

		cities, err := servicesRO.GetCityService()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		destinationCity := strings.Title(dataCheckOut.Address.City)
		deliveryCost, errCost := servicesRO.GetDeliveryCostService(cities[destinationCity])
		if errCost != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err,
			})
		}

		userID, _ := strconv.Atoi(c.Param("userID"))
		product, _ := database.GetProductById(strconv.Itoa(int(dataCheckOut.ProductID)))
		totalProductPrice := product.Price * dataCheckOut.QTY
		totalPrice := totalProductPrice + uint(deliveryCost)
		transactionNumber := utils.GenerateRandomTransactionID()
		userEmail := database.GetUserEmail(c.Param("userID"))

		midtransRequest := models.MidtransRequest{TransactionNumber: transactionNumber, Amount: int64(totalPrice),
			Product: models.AllProductResponse{
				ID:   dataCheckOut.ProductID,
				Name: product.Name, Price: product.Price},
			QTY:          int32(dataCheckOut.QTY),
			ShippingCost: int64(deliveryCost),
			User: struct {
				Name  string
				Email string
				Phone string
			}{
				dataCheckOut.Name,
				userEmail,
				dataCheckOut.PhoneNumber}}

		checkOutResponse, err := midtrans.CreateSnapToken(&midtransRequest)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err,
			})
		}

		product.Stock = product.Stock - dataCheckOut.QTY
		err = database.SaveProduct(&product)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err,
			})
		}

		// Save shipping
		address := dataCheckOut.Address.Detail + ", " + dataCheckOut.Address.District + ", " +
			dataCheckOut.Address.City + ", " + dataCheckOut.Address.Province + ", " + dataCheckOut.Address.PostalCode
		shipping.Name = dataCheckOut.Name
		shipping.Address = address
		shipping.PhoneNumber = dataCheckOut.PhoneNumber
		err = database.SaveShipping(&shipping)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": err,
			})
		}

		// Save transaction
		transaction.TransactionNumber = transactionNumber
		transaction.Date = time.Now().Format("2006-01-02 15:04:05")
		transaction.UserID = uint(userID)
		transaction.TotalQTY = dataCheckOut.QTY
		transaction.ShippingCost = uint(deliveryCost)
		transaction.TotalPrice = totalPrice
		transaction.PaymentStatus = "pending"
		transaction.SnapToken = checkOutResponse.Token
		errTransaction := database.SaveTransaction(&transaction)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": errTransaction.Error(),
			})
		}

		// Save transaction detail
		var transactionDetail models.TransactionDetail
		transactionDetail.TransactionID = transaction.ID
		transactionDetail.ProductID = dataCheckOut.ProductID
		transactionDetail.QTY = dataCheckOut.QTY
		transactionDetail.Price = product.Price
		transactionDetail.ShippingID = shipping.ID
		errTD := database.SaveTransactionDetail(&transactionDetail)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": errTD.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message":  "Success",
			"ID":       transaction.ID,
			"midtrans": checkOutResponse,
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
		transactionsResponse.PaymentMethod = transactions[i].PaymentMethod
		transactionsResponse.PaymentStatus = transactions[i].PaymentStatus
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

	if transactionDetail.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Not found",
		})
	}

	transactionDetailResponse := models.TransactionDetailResponse{ID: transactionDetail.ID,
		TransactionID: transactionDetail.TransactionID, ProductID: transactionDetail.ProductID,
		QTY: transactionDetail.QTY, Price: transactionDetail.Price,
		Transaction: struct {
			TransactionNumber string
			Date              string
			PaymentStatus     string
		}{transactionDetail.Transaction.TransactionNumber,
			transactionDetail.Transaction.Date,
			transactionDetail.Transaction.PaymentStatus},
		Product: models.AllProductResponse{
			ID:       transactionDetail.Product.ID,
			Name:     transactionDetail.Product.Name,
			Price:    transactionDetail.Product.Price,
			ImageURI: transactionDetail.Product.Image_URI},
		Address: transactionDetail.Shipping}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":      "Success Get Transaction Detail",
		"transactions": transactionDetailResponse,
	})
}
