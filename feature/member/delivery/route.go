package delivery

import (
	"wilayah/config"
	"wilayah/domain"
	"wilayah/feature/member/delivery/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouteMember(e *echo.Echo, member domain.MemberHandler) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	e.POST("/member", member.AddMember(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.PUT("/member/:id", member.UpdateDataMember(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.DELETE("/member/:id", member.DeleteDataMember(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
	e.GET("/member", member.GetMember(), middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET))))
}
