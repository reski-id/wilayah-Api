package delivery

import (
	"wilayah/config"
	"wilayah/domain"
	"wilayah/feature/kelurahan/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteKel(e *echo.Echo, kelurahan domain.KelurahanHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/kelurahan", kelurahan.AddKelurahan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/kelurahan/:id", kelurahan.UpdateDataKelurahan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/kelurahan/:id", kelurahan.DeleteDataKelurahan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/kelurahan", kelurahan.GetKelurahan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
