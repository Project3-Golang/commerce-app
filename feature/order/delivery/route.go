package delivery

import (
	"commerce-app/config"
	"commerce-app/domain"
	"commerce-app/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteOrder(e *echo.Echo, oh domain.OrderHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.POST("/order", oh.InsertOrder())
	e.PUT("/order/:id", oh.UpdateOrder(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/order/:id", oh.DeleteOrder(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/order", oh.GetAllOrder())
	e.GET("/order/:id", oh.GetOrderID())
}
