package domain

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Order struct {
	ID        int
	User      string
	Cart      int
	Status    int
	Payment   string
	Total     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderUseCase interface {
	AddOrder(newOrder Order) (Order, error)
	GetAllO() ([]Order, error)
	UpOrder(IDOrder int, updateData Order) (Order, error)
	DelOrder(IDOrder int) (bool, error)
	GetSpecificOrder(orderID int) ([]Order, error)
}

type OrderHandler interface {
	InsertOrder() echo.HandlerFunc
	GetAllOrder() echo.HandlerFunc
	UpdateOrder() echo.HandlerFunc
	DeleteOrder() echo.HandlerFunc
	GetOrderID() echo.HandlerFunc
}

type OrderData interface {
	Insert(insertOrder Order) Order
	GetAll() []Order
	Update(IDOrder int, updatedOrder Order) Order
	Delete(IDOrder int) bool
	GetOrderID(orderID int) []Order
}
