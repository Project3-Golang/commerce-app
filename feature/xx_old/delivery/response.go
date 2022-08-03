package delivery

import "commerce-app/domain"

type CartResponse struct {
	ID        int  `json:"id"`
	Quantity  int  `json:"qty"`
	UserID    uint `json:"user_id"`
	ProductID uint `json:"product_id"`
}

func FromDomain(data domain.Cart) CartResponse {
	var res CartResponse
	res.ID = data.ID
	res.UserID = uint(data.UserID)
	res.Quantity = data.Quantity
	return res
}
