package delivery

import "commerce-app/domain"

type ProductResponse struct {
	ID          int    `json:"id"`
	Name        string `json:"name" form:"name"`
	Price       int32  `json:"price" form:"price"`
	Stock       int32  `json:"stock" form:"stock"`
	Description string `json:"description" form:"description"`
	Images      string `json:"images" form:"images"`
}

func FromDomain(data domain.Product) ProductResponse {
	var res ProductResponse
	res.ID = int(data.ID)
	res.Name = data.Name
	res.Price = int32(data.Price)
	res.Stock = int32(data.Stock)
	res.Description = data.Description
	res.Images = data.Images
	return res
}
