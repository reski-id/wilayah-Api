package usecase

import (
	"errors"
	"wilayah/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type kotaUseCase struct {
	kotaData domain.KotaData
	validate *validator.Validate
}

func New(kd domain.KotaData, v *validator.Validate) domain.KotaUsecase {
	return &kotaUseCase{
		kotaData: kd,
		validate: v,
	}
}

func (ku *kotaUseCase) AddKota(newData domain.Kota) (idKota int, err error) {
	if newData.KotaName == "" {
		return -1, errors.New("invalid")
	}
	inserted, err := ku.kotaData.InsertKota(newData)
	return inserted, err
}

func (ku *kotaUseCase) GetAllKota() ([]domain.Kota, error) {
	data, err := ku.kotaData.GetAll()
	return data, err
}

func (ku *kotaUseCase) UpdateKota(id int, UpdateData domain.Kota) (idKota int, err error) {

	if UpdateData.KotaName == "" {
		return -1, errors.New("invalid")
	}

	data, err := ku.kotaData.InsertKota(UpdateData)
	return data, err
}

func (ku *kotaUseCase) DeleteKota(id int) (idKota int, err error) {
	idKota, err = ku.kotaData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idKota, errors.New("data not found")
		} else {
			return idKota, errors.New("failed to delete data")
		}
	}
	return idKota, nil
}
