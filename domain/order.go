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
}

type OrderHandler interface {
	InsertOrder() echo.HandlerFunc
}

type OrderData interface {
	Insert(insertOrder Order) Order
}
