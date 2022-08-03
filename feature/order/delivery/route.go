package delivery

import (
	"commerce-app/domain"

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
	// e.PUT("/product/:id", bc.UpdateProduct(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	// e.DELETE("/product/:id", bc.DeleteProduct(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	// e.GET("/product", bc.GetAllProduct())
	// e.GET("/product/:id", bc.GetProductID())
}
