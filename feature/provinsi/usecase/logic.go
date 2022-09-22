package usecase

import (
	"errors"
	"wilayah/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type provinsiUseCase struct {
	provinsiData domain.ProvinsiData
	validate     *validator.Validate
}

func New(pd domain.ProvinsiData, v *validator.Validate) domain.ProvinsiUsecase {
	return &provinsiUseCase{
		provinsiData: pd,
		validate:     v,
	}
}

func (pu *provinsiUseCase) AddProvinsi(newData domain.Provinsi) (idProvinsi int, err error) {
	if newData.ProvinsiName == "" {
		return -1, errors.New("invalid")
	}
	inserted, err := pu.provinsiData.InsertProvinsi(newData)
	return inserted, err
}

func (pu *provinsiUseCase) GetAllProvinsi() ([]domain.Provinsi, error) {
	data, err := pu.provinsiData.GetAll()
	return data, err
}

func (pu *provinsiUseCase) UpdateProvinsi(id int, UpdateData domain.Provinsi) (idProvinsi int, err error) {

	if UpdateData.ProvinsiName == "" {
		return -1, errors.New("invalid")
	}

	data, err := pu.provinsiData.InsertProvinsi(UpdateData)
	return data, err
}

func (pu *provinsiUseCase) DeleteProvinsi(id int) (idProvinsi int, err error) {
	idProvinsi, err = pu.provinsiData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idProvinsi, errors.New("data not found")
		} else {
			return idProvinsi, errors.New("failed to delete data")
		}
	}
	return idProvinsi, nil
}
