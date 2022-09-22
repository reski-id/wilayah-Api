package factory

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	ud "wilayah/feature/user/data"
	userDelivery "wilayah/feature/user/delivery"
	us "wilayah/feature/user/usecase"

	pd "wilayah/feature/provinsi/data"
	provinsiDelivery "wilayah/feature/provinsi/delivery"
	pu "wilayah/feature/provinsi/usecase"

	kd "wilayah/feature/kota/data"
	kotaDelivery "wilayah/feature/kota/delivery"
	ku "wilayah/feature/kota/usecase"

	kecd "wilayah/feature/kecamatan/data"
	kecDelivery "wilayah/feature/kecamatan/delivery"
	kecu "wilayah/feature/kecamatan/usecase"
)

func Initfactory(e *echo.Echo, db *gorm.DB) {
	//factory
	userData := ud.New(db)
	validator := validator.New()
	useCase := us.New(userData, validator)
	userDelivery.New(e, useCase)

	provinsiData := pd.New(db)
	provinsiCase := pu.New(provinsiData, validator)
	provinsiHandler := provinsiDelivery.New(provinsiCase)
	provinsiDelivery.RouteProvinsi(e, provinsiHandler)

	kotaData := kd.New(db)
	kotaCase := ku.New(kotaData, validator)
	kotaHandler := kotaDelivery.New(kotaCase)
	kotaDelivery.RouteKota(e, kotaHandler)

	kecData := kecd.New(db)
	kecCase := kecu.New(kecData, validator)
	kecHandler := kecDelivery.New(kecCase)
	kecDelivery.RouteKec(e, kecHandler)
}
