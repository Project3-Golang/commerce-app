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

	cd "commerce-app/feature/cart/data"
	cartDelivery "commerce-app/feature/cart/delivery"
	cu "commerce-app/feature/cart/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	productData := pd.New(db)
	productCase := pu.New(productData, validator)
	productHandler := productDelivery.New(productCase)
	productDelivery.RouteProduct(e, productHandler)

	cartData := cd.New(db)
	cartCase := cu.New(cartData, validator)
	cartHandler := cartDelivery.New(cartCase)
	cartDelivery.RouteCart(e, cartHandler)
}
