package domain

import "github.com/labstack/echo/v4"

type Kota struct {
	ID           int
	KotaName     string
	ProvinsiCode string
}

type KotaHandler interface {
	AddKota() echo.HandlerFunc
	GetKota() echo.HandlerFunc
	UpdateDataKota() echo.HandlerFunc
	DeleteDataKota() echo.HandlerFunc
}

type KotaUsecase interface {
	AddKota(newKota Kota) (row int, err error)
	GetAllKota() ([]Kota, error)
	UpdateKota(id int, UpdateKota Kota) (row int, err error)
	DeleteKota(id int) (row int, err error)
}

type KotaData interface {
	InsertKota(newKota Kota) (row int, err error)
	GetAll() ([]Kota, error)
	Update(id int, updatedData Kota) (row int, err error)
	Delete(id int) (row int, err error)
}
