package delivery

import "commerce-app/domain"

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" form:"name"`
	Price       float64 `json:"price" form:"price"`
	Stock       int32   `json:"stock" form:"stock"`
	Description string  `json:"desription" form:"desription"`
	Images      string  `json:"images_url" form:"images_url"`
}

func FromDomain(data domain.Product) ProductResponse {
	var res ProductResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Images = data.Images
	res.Price = data.Price
	res.Stock = int32(data.Stock)
	return res
}
