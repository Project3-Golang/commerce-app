package delivery

import (
	"commerce-app/domain"
	"time"
)

type ProductInsertRequest struct {
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Stock       int32   `json:"stock" form:"stock"`
	Description string  `json:"desription" form:"desription"`
	Images      string  `json:"images_url" form:"images_url"`
}

func (pi *ProductInsertRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:        pi.Name,
		Price:       pi.Price,
		Stock:       int(pi.Stock),
		Description: pi.Description,
		Images:      pi.Images,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
