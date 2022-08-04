package data

import (
	"commerce-app/domain"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	UserID    uint   `json:"user_id" form:"user_id"`
	ProductID uint   `json:"product_id" form:"product_id"`
	CartID    uint   `json:"cart_id" form:"cart_id"`
	Status    int    `json:"status" form:"status"   gorm:"default:0"`
	Payment   string `json:"payment" form:"payment"`
	Total     string `json:"total" form:"total"`
}
type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images"`
}
type Cart struct {
	gorm.Model
	Quantity int `json:"quantity" form:"quantity" validate:"required"`
}
type User struct {
	gorm.Model
	Username string `json:"username" form:"username" validate:"required"`
	Email    string `gorm:"unique" json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Role     string `json:"role" form:"role" gorm:"default:users"`
	Photo    string `json:"image_url"`
}

func (o *Order) ToDomain() domain.Order {
	return domain.Order{
		ID:        int(o.ID),
		CartID:    int(o.CartID),
		ProductID: int(o.ProductID),
		UserID:    int(o.UserID),
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
	res.UserID = uint(data.UserID)
	res.CartID = uint(data.CartID)
	res.ProductID = uint(data.ProductID)
	res.Status = data.Status
	res.Payment = data.Payment
	res.Total = data.Total
	return res
}
