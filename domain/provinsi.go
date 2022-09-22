package domain

import "github.com/labstack/echo/v4"

type Provinsi struct {
	ID           int
	ProvinsiName string
	ProvinsiCode string
}

type ProvinsiHandler interface {
	AddProvinsi() echo.HandlerFunc
	GetProvinsi() echo.HandlerFunc
	UpdateDataProvinsi() echo.HandlerFunc
	DeleteDataProvinsi() echo.HandlerFunc
}

type ProvinsiUsecase interface {
	AddProvinsi(newProvinsi Provinsi) (row int, err error)
	GetAllProvinsi() ([]Provinsi, error)
	UpdateProvinsi(id int, UpdateProvinsi Provinsi) (row int, err error)
	DeleteProvinsi(id int) (row int, err error)
}

type ProvinsiData interface {
	InsertProvinsi(newProvinsi Provinsi) (row int, err error)
	GetAll() ([]Provinsi, error)
	Update(id int, updatedData Provinsi) (row int, err error)
	Delete(id int) (row int, err error)
}
