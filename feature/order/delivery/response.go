package delivery

import "commerce-app/domain"

type OrderResponse struct {
	ID      int    `json:"id"`
	User    string `json:"user" form:"user"`
	Cart    int    `json:"cart" form:"cart"`
	Status  int    `json:"status" form:"status"`
	Payment string `json:"payment" form:"payment"`
	Total   string `json:"total" form:"total"`
}

func FromDomain(data domain.Order) OrderResponse {
	var res OrderResponse
	res.ID = int(data.ID)
	res.Cart = data.Cart
	res.Payment = data.Payment
	res.Status = data.Status
	res.Total = data.Total
	return res
}
