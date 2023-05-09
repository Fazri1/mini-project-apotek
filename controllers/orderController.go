package controllers

import (
	"fmt"
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
		var order models.Order
		c.Bind(&dataCheckOut)

		cities, err := servicesRO.GetCityService()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}
		fmt.Println("ok")
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
		orderNumber := utils.GenerateRandomOrderID()
		userEmail := database.GetUserEmail(c.Param("userID"))

		midtransRequest := models.MidtransRequest{OrderNumber: orderNumber, Amount: int64(totalPrice),
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

		// Save order
		order.OrderNumber = orderNumber
		order.Date = time.Now().Format("2006-01-02 15:04:05")
		order.UserID = uint(userID)
		order.TotalQTY = dataCheckOut.QTY
		order.ShippingCost = uint(deliveryCost)
		order.TotalPrice = totalPrice
		order.Status = "waiting payment"
		order.PaymentStatus = "pending"
		order.SnapToken = checkOutResponse.Token
		errOrder := database.SaveOrder(&order)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": errOrder.Error(),
			})
		}

		// Save order detail
		var orderDetail models.OrderDetail
		orderDetail.OrderID = order.ID
		orderDetail.ProductID = dataCheckOut.ProductID
		orderDetail.QTY = dataCheckOut.QTY
		orderDetail.Price = product.Price
		orderDetail.ShippingID = shipping.ID
		errTD := database.SaveOrderDetail(&orderDetail)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": errTD.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message":  "Success",
			"ID":       order.ID,
			"midtrans": checkOutResponse,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}

func GetUserOrdersController(c echo.Context) error {
	orders, err := database.GetUserOrders(c.Param("userID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	ordersList := []models.OrderResponse{}
	for i := range orders {
		orderResponse := models.OrderResponse{}
		orderResponse.ID = orders[i].ID
		orderResponse.OrderNumber = orders[i].OrderNumber
		orderResponse.Date = orders[i].Date
		orderResponse.UserID = orders[i].UserID
		orderResponse.TotalQTY = orders[i].TotalQTY
		orderResponse.ShippingCost = orders[i].ShippingCost
		orderResponse.TotalPrice = orders[i].TotalPrice
		orderResponse.Status = orders[i].Status
		orderResponse.PaymentMethod = orders[i].PaymentMethod
		orderResponse.PaymentStatus = orders[i].PaymentStatus
		orderResponse.User.ID = orders[i].User.ID
		orderResponse.User.Name = orders[i].User.Name
		orderResponse.User.Email = orders[i].User.Email
		ordersList = append(ordersList, orderResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get orders",
		"orders":  ordersList,
	})
}

func GetUserOrderDetailController(c echo.Context) error {
	orderDetail, err := database.GetUserOrderDetail(c.Param("userID"), c.Param("orderID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": err.Error(),
		})
	}

	if orderDetail.ID == 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "Not found",
		})
	}

	orderDetailResponse := models.OrderDetailResponse{ID: orderDetail.ID,
		OrderID: orderDetail.OrderID, ProductID: orderDetail.ProductID,
		QTY: orderDetail.QTY, Price: orderDetail.Price,
		Order: struct {
			OrderNumber   string
			Date          string
			Status        string
			PaymentStatus string
		}{orderDetail.Order.OrderNumber,
			orderDetail.Order.Date,
			orderDetail.Order.Status,
			orderDetail.Order.PaymentStatus},
		Product: models.AllProductResponse{
			ID:       orderDetail.Product.ID,
			Name:     orderDetail.Product.Name,
			Price:    orderDetail.Product.Price,
			ImageURI: orderDetail.Product.ImageURI},
		Address: orderDetail.Shipping}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success get order detail",
		"orders":  orderDetailResponse,
	})
}

func GetAllOrdersController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		orders, err := database.GetAllOrders()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"message": err.Error(),
			})
		}

		ordersList := []models.OrderResponse{}
		for i := range orders {
			orderResponse := models.OrderResponse{}
			orderResponse.ID = orders[i].ID
			orderResponse.OrderNumber = orders[i].OrderNumber
			orderResponse.Date = orders[i].Date
			orderResponse.UserID = orders[i].UserID
			orderResponse.TotalQTY = orders[i].TotalQTY
			orderResponse.ShippingCost = orders[i].ShippingCost
			orderResponse.TotalPrice = orders[i].TotalPrice
			orderResponse.Status = orders[i].Status
			orderResponse.PaymentMethod = orders[i].PaymentMethod
			orderResponse.PaymentStatus = orders[i].PaymentStatus
			orderResponse.User.ID = orders[i].User.ID
			orderResponse.User.Name = orders[i].User.Name
			orderResponse.User.Email = orders[i].User.Email
			ordersList = append(ordersList, orderResponse)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get all orders",
			"orders":  ordersList,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}

func GetOrderDetailController(c echo.Context) error {
	token := strings.Fields(c.Request().Header.Values("Authorization")[0])[1]
	admin, err := middlewares.CheckTokenRole(token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	if admin {
		orderDetail, err := database.GetOrderDetail(c.Param("orderID"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"message": err.Error(),
			})
		}

		orderDetailResponse := models.OrderDetailResponse{ID: orderDetail.ID,
			OrderID: orderDetail.OrderID, ProductID: orderDetail.ProductID,
			QTY: orderDetail.QTY, Price: orderDetail.Price,
			Order: struct {
				OrderNumber   string
				Date          string
				Status        string
				PaymentStatus string
			}{orderDetail.Order.OrderNumber,
				orderDetail.Order.Date,
				orderDetail.Order.Status,
				orderDetail.Order.PaymentStatus},
			Product: models.AllProductResponse{
				ID:       orderDetail.Product.ID,
				Name:     orderDetail.Product.Name,
				Price:    orderDetail.Product.Price,
				ImageURI: orderDetail.Product.ImageURI},
			Address: orderDetail.Shipping}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Success get order detail",
			"order":   orderDetailResponse,
		})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{
		"message": "Unauthorized action",
	})
}

func UpdateStatusOrderController(c echo.Context) error {
	var order models.Order
	c.Bind(&order)

	err := database.UpdateStatusOrder(order.Status, c.Param("orderID"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success update status",
	})
}
