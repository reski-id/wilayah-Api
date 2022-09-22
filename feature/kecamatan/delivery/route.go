package delivery

import (
	"wilayah/config"
	"wilayah/domain"
	"wilayah/feature/kecamatan/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteKec(e *echo.Echo, kecamatan domain.KecamatanHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/kecamatan", kecamatan.AddKecamatan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/kecamatan/:id", kecamatan.UpdateDataKecamatan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/kecamatan/:id", kecamatan.DeleteDataKecamatan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/kecamatan", kecamatan.GetKecamatan(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
