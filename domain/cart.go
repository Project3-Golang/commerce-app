package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Cart struct {
	ID        int
	Quantity  int
	ProductID int
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CartUseCase interface {
	AddCart(IDUser int, useCart Cart) (Cart, error)
	GetAllC() ([]Cart, error)
	UpCart(IDCart int, updateData Cart) (Cart, error)
	DelCart(IDCart int) (bool, error)
	GetSpecificCart(cartID int) ([]Cart, error)
	GetmyCart(userID int) ([]Cart, error)
}

type CartHandler interface {
	InsertCart() echo.HandlerFunc
	GetAllCart() echo.HandlerFunc
	UpdateCart() echo.HandlerFunc
	DeleteCart() echo.HandlerFunc
	GetCartID() echo.HandlerFunc
	GetMYCart() echo.HandlerFunc
}

type CartData interface {
	Insert(insertCart Cart) Cart
	GetAll() []Cart
	Update(IDCart int, updatedCart Cart) Cart
	Delete(IDCart int) bool
	GetCartID(cartID int) []Cart
	GetCartbyuser(userID int) []Cart
}
