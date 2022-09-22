package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "wilayah/feature/user/data"
	userDelivery "wilayah/feature/user/delivery"
	us "wilayah/feature/user/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	//factory
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)
}
