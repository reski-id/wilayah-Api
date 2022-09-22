package usecase

import (
	"errors"
	"wilayah/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type kelUseCase struct {
	kelData  domain.KelurahanData
	validate *validator.Validate
}

func New(kd domain.KelurahanData, v *validator.Validate) domain.KelurahanUsecase {
	return &kelUseCase{
		kelData:  kd,
		validate: v,
	}
}

func (ku *kelUseCase) AddKelurahan(newData domain.Kelurahan) (idKelurahan int, err error) {
	if newData.KelurahanName == "" {
		return -1, errors.New("invalid")
	}
	inserted, err := ku.kelData.InsertKelurahan(newData)
	return inserted, err
}

func (ku *kelUseCase) GetAllKelurahan() ([]domain.Kelurahan, error) {
	data, err := ku.kelData.GetAll()
	return data, err
}

func (ku *kelUseCase) UpdateKelurahan(id int, UpdateData domain.Kelurahan) (idKelurahan int, err error) {

	if UpdateData.KelurahanName == "" {
		return -1, errors.New("invalid")
	}

	data, err := ku.kelData.InsertKelurahan(UpdateData)
	return data, err
}

func (ku *kelUseCase) DeleteKelurahan(id int) (idKelurahan int, err error) {
	idKelurahan, err = ku.kelData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idKelurahan, errors.New("data not found")
		} else {
			return idKelurahan, errors.New("failed to delete data")
		}
	}
	return idKelurahan, nil
}
