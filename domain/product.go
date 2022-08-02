package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int
	Name        string
	Price       float64
	Stock       int
	Description string
	Images      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductUseCase interface {
	AddProduct(newProduct Product) (Product, error)
	GetProduct(id int) (Product, error)
	DeleteProduct(id int) (bool, error)
	UpdateProduct(id int, updateProduct Product) (Product, error)
	GetAll() ([]Product, error)
}

type ProductData interface {
	Insert(newProduct Product) (Product, error)
	GetSpecific(productID int) (Product, error)
	Delete(productID int) bool
	Update(productID int, updatedData Product) Product
	GetAll() []Product
}

type ProductHandler interface {
	InsertProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
	GetProductID() echo.HandlerFunc
}
