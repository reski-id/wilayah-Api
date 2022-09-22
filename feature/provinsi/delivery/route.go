package delivery

import (
	"wilayah/config"
	"wilayah/domain"
	"wilayah/feature/provinsi/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteProvinsi(e *echo.Echo, provinsi domain.ProvinsiHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/provinsi", provinsi.AddProvinsi(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/provinsi/:id", provinsi.UpdateDataProvinsi(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/provinsi/:id", provinsi.DeleteDataProvinsi(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/provinsi", provinsi.GetProvinsi(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
