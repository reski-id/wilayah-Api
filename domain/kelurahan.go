package domain

import "github.com/labstack/echo/v4"

type Kelurahan struct {
	ID            int
	KelurahanName string
	KecamatanID   string
}

type KelurahanHandler interface {
	AddKelurahan() echo.HandlerFunc
	GetKelurahan() echo.HandlerFunc
	UpdateDataKelurahan() echo.HandlerFunc
	DeleteDataKelurahan() echo.HandlerFunc
}

type KelurahanUsecase interface {
	AddKelurahan(newKelurahan Kelurahan) (row int, err error)
	GetAllKelurahan() ([]Kelurahan, error)
	UpdateKelurahan(id int, UpdateKelurahan Kelurahan) (row int, err error)
	DeleteKelurahan(id int) (row int, err error)
}

type KelurahanData interface {
	InsertKelurahan(newKelurahan Kelurahan) (row int, err error)
	GetAll() ([]Kelurahan, error)
	Update(id int, updatedData Kelurahan) (row int, err error)
	Delete(id int) (row int, err error)
}
