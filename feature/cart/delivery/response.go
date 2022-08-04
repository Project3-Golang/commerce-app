package delivery

import "commerce-app/domain"

type CartResponse struct {
	ID        int  `json:"id"`
	Quantity  int  `json:"quantity"`
	UserID    int  `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func FromDomain(data domain.Cart) CartResponse {
	var res CartResponse
	res.ID = data.ID
	res.UserID = data.UserID
	res.Quantity = data.Quantity
	res.ProductID = uint(data.ProductID)
	return res
}
