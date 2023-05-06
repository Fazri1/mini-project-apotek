package midtrans

import (
	"mini-project-apotek/constants"
	"mini-project-apotek/models"
	"strconv"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

func CreateSnapToken(request *models.MidtransRequest) (*snap.Response, error) {
	var s = snap.Client{}
	s.New(constants.MT_SERVER_KEY, midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  request.TransactionNumber,
			GrossAmt: request.Amount,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.Itoa(int(request.Product.ID)),
				Name:  request.Product.Name,
				Price: int64(request.Product.Price),
				Qty:   request.QTY,
			},
			{
				Name:  "Ongkos Kirim",
				Price: request.ShippingCost,
				Qty:   1,
			},
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: request.User.Name,
			Email: request.User.Email,
			Phone: request.User.Phone,
		},
	}
	snapResp, _ := s.CreateTransaction(req)
	// if err != nil {
	// 	return &snap.Response{}, err
	// }
	return snapResp, nil
}
