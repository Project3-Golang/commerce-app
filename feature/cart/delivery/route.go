package delivery

import (
	"commerce-app/config"
	"commerce-app/domain"
	"commerce-app/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteCart(e *echo.Echo, bc domain.CartHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/cart", bc.InsertCart(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/cart/:id", bc.UpdateCart(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/cart/:id", bc.DeleteCart(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/cart", bc.GetAllCart())
	e.GET("/mycart", bc.GetMYCart())
	e.GET("/cart/:id", bc.GetCartID())
}
