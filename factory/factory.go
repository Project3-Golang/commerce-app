package factory

import (
	ud "commerce-app/feature/user/data"
	userDelivery "commerce-app/feature/user/delivery"
	us "commerce-app/feature/user/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	pd "commerce-app/feature/product/data"
	productDelivery "commerce-app/feature/product/delivery"
	pu "commerce-app/feature/product/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	productData := pd.New(db)
	productCase := pu.New(productData, validator)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteBook(e, productHandler)
}
