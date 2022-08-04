package data

import (
	"commerce-app/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	User    string `json:"user" form:"user" validate:"required"`
	Cart    int    `json:"cart" form:"cart" validate:"required"`
	Status  int    `json:"status" form:"status" validate:"required"`
	Payment string `json:"payment" form:"payment"`
	Total   string `json:"total" form:"total"`
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:        int(o.ID),
		User:      o.User,
		Cart:      o.Cart,
		Status:    o.Status,
		Payment:   o.Payment,
		Total:     o.Total,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}

func ParseToArr(arr []Order) []domain.Order {
	var res []domain.Order

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Order) Order {
	var res Order
	res.User = data.User
	res.Cart = data.Cart
	res.Status = data.Status
	res.Payment = data.Payment
	res.Total = data.Total
	return res
}
