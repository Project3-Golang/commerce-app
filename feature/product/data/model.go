package data

import (
	"commerce-app/domain"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string  `json:"name" form:"name" validate:"required"`
	Price       float64 `json:"price" form:"price" validate:"required"`
	Stock       int     `json:"stock" form:"stock" validate:"required"`
	Description string  `json:"description" form:"description"`
	Images      string  `json:"images"`
}

func (p *Product) ToDomain() domain.Product {
	return domain.Product{
		ID:          int(p.ID),
		Name:        p.Name,
		Price:       p.Price,
		Stock:       p.Stock,
		Description: p.Description,
		Images:      p.Images,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
	}
}

func ParseToArr(arr []Product) []domain.Product {
	var res []domain.Product

	for _, val := range arr {
		res = append(res, val.ToDomain())
	}
	return res
}

func ToLocal(data domain.Product) Product {
	var res Product
	res.Name = data.Name
	res.Price = data.Price
	res.Stock = data.Stock
	res.Description = data.Description
	res.Images = data.Images
	return res
}
