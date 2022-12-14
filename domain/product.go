package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID          int
	Name        string
	Price       int
	Stock       int
	Description string
	Images      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type ProductUseCase interface {
	AddProduct(newProduct Product) (Product, error)
	GetAllP() ([]Product, error)
	UpProduct(IDProduct int, updateData Product) (Product, error)
	DelProduct(IDProduct int) (bool, error)
	GetSpecificProduct(productID int) ([]Product, error)
}

type ProductHandler interface {
	InsertProduct() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetProductID() echo.HandlerFunc
}

type ProductData interface {
	Insert(insertProduct Product) Product
	GetAll() []Product
	Update(IDProduct int, updatedProduct Product) Product
	Delete(IDProduct int) bool
	GetProductID(productID int) []Product
}
