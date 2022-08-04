package delivery

import (
	"commerce-app/domain"
	"time"
)

type OrderInsertRequest struct {
	UserID    uint   `json:"user_id" form:"user_id"`
	ProductID uint   `json:"product_id" form:"product_id"`
	CartID    uint   `json:"cart_id" form:"cart_id"`
	Status    int    `json:"status" form:"status"`
	Payment   string `json:"payment" form:"payment"`
	Total     string `json:"total" form:"total"`
}

func (oi *OrderInsertRequest) ToDomain() domain.Order {
	return domain.Order{
		CartID:    int(oi.CartID),
		ProductID: int(oi.ProductID),
		UserID:    int(oi.UserID),
		Status:    oi.Status,
		Payment:   oi.Payment,
		Total:     oi.Total,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
