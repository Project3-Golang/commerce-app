package delivery

import (
	"commerce-app/domain"
	"time"
)

type CartInsertRequest struct {
	Quantity  int  `json:"quantity" form:"quantity" validate:"required"`
	UserID    uint `json:"user_id" form:"user_id"`
	ProductID uint `json:"product_id" form:"product_id"`
}

func (ci *CartInsertRequest) ToDomain() domain.Cart {
	return domain.Cart{
		Quantity:  ci.Quantity,
		ProductID: int(ci.ProductID),
		UserID:    int(ci.UserID),
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
