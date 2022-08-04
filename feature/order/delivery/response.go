package delivery

import "commerce-app/domain"

type OrderResponse struct {
	ID        int    `json:"id"`
	UserID    uint   `json:"user_id" form:"user_id"`
	ProductID uint   `json:"product_id" form:"product_id"`
	CartID    uint   `json:"cart_id" form:"cart_id"`
	Status    int    `json:"status" form:"status"`
	Payment   string `json:"payment" form:"payment"`
	Total     string `json:"total" form:"total"`
}

func FromDomain(data domain.Order) OrderResponse {
	var res OrderResponse
	res.ID = int(data.ID)
	res.CartID = uint(data.CartID)
	res.ProductID = uint(data.ProductID)
	res.UserID = uint(data.UserID)
	res.Payment = data.Payment
	res.Status = data.Status
	res.Total = data.Total
	return res
}
