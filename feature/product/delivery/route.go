package delivery

import (
	"commerce-app/config"
	"commerce-app/domain"
	"commerce-app/feature/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteProduct(e *echo.Echo, bc domain.ProductHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	// e.POST("/product", bc.InsertProduct(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.PUT("/product/:id", bc.UpdateProduct(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.DELETE("/product/:id", bc.DeleteProduct(), middleware.JWTWithConfig(common.UseJWT([]byte(config.SECRET))))
	e.GET("/product", bc.GetAllProduct())
	e.GET("/product/:id", bc.GetProductID())
}
