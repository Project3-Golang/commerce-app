package data

import (
	"commerce-app/domain"

	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Quantity  int  `json:"qty" form:"qty" validate:"required"`
	UserID    uint `json:"user_id" form:"user_id"`
	ProductID uint `json:"product_id" form:"product_id"`
}

type Product struct {
	gorm.Model
	Name        string `json:"name" form:"name" validate:"required"`
	Price       int    `json:"price" form:"price" validate:"required"`
	Stock       int    `json:"stock" form:"stock" validate:"required"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images"`
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

func (c *Cart) ToDomain() domain.Cart {
	return domain.Cart{
		ID:        int(c.ID),
		Quantity:  int(c.Quantity),
		ProductID: int(c.ProductID),
		UserID:    int(c.UserID),
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func ParseToArr(arr []Cart) []domain.Cart {
	var res []domain.Cart

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Cart) Cart {
	var res Cart
	res.ID = uint(data.ID)
	res.ProductID = uint(data.ProductID)
	res.Quantity = data.Quantity
	res.UserID = uint(data.UserID)

	return res
}
