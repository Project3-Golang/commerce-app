package delivery

import (
	"commerce-app/domain"
	"time"
)

type ProductInsertRequest struct {
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Stock       int32   `json:"stock" form:"stock"`
	Description string  `json:"description" form:"description"`
	Images      string  `json:"images" form:"images"`
}

func (pi *ProductInsertRequest) ToDomain() domain.Product {
	return domain.Product{
		Name:        pi.Name,
		Price:       int(pi.Price),
		Stock:       int(pi.Stock),
		Description: pi.Description,
		Images:      pi.Images,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
}
