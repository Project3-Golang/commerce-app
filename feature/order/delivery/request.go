package delivery

import (
	"commerce-app/domain"
	"time"
)

type OrderInsertRequest struct {
	User    string `json:"user" form:"user"`
	Cart    int    `json:"cart" form:"cart"`
	Status  int    `json:"status" form:"status"`
	Payment string `json:"payment" form:"payment"`
	Total   string `json:"total" form:"total"`
}

func (oi *OrderInsertRequest) ToDomain() domain.Order {
	return domain.Order{
		User:      oi.User,
		Cart:      oi.Cart,
		Status:    oi.Status,
		Payment:   oi.Payment,
		Total:     oi.Total,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
