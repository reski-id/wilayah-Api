package domain

import "github.com/labstack/echo/v4"

type Kecamatan struct {
	ID            int
	KecamatanName string
	KotaID        string
}

type KecamatanHandler interface {
	AddKecamatan() echo.HandlerFunc
	GetKecamatan() echo.HandlerFunc
	UpdateDataKecamatan() echo.HandlerFunc
	DeleteDataKecamatan() echo.HandlerFunc
}

type KecamatanUsecase interface {
	AddKecamatan(newKecamatan Kecamatan) (row int, err error)
	GetAllKecamatan() ([]Kecamatan, error)
	UpdateKecamatan(id int, UpdateKecamatan Kecamatan) (row int, err error)
	DeleteKecamatan(id int) (row int, err error)
}

type KecamatanData interface {
	InsertKecamatan(newKecamatan Kecamatan) (row int, err error)
	GetAll() ([]Kecamatan, error)
	Update(id int, updatedData Kecamatan) (row int, err error)
	Delete(id int) (row int, err error)
}
