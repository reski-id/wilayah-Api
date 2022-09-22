package usecase

import (
	"errors"
	"wilayah/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type kecUseCase struct {
	kecData  domain.KecamatanData
	validate *validator.Validate
}

func New(kd domain.KecamatanData, v *validator.Validate) domain.KecamatanUsecase {
	return &kecUseCase{
		kecData:  kd,
		validate: v,
	}
}

func (ku *kecUseCase) AddKecamatan(newData domain.Kecamatan) (idKecamatan int, err error) {
	if newData.KecamatanName == "" {
		return -1, errors.New("invalid")
	}
	inserted, err := ku.kecData.InsertKecamatan(newData)
	return inserted, err
}

func (ku *kecUseCase) GetAllKecamatan() ([]domain.Kecamatan, error) {
	data, err := ku.kecData.GetAll()
	return data, err
}

func (ku *kecUseCase) UpdateKecamatan(id int, UpdateData domain.Kecamatan) (idKecamatan int, err error) {

	if UpdateData.KecamatanName == "" {
		return -1, errors.New("invalid")
	}

	data, err := ku.kecData.InsertKecamatan(UpdateData)
	return data, err
}

func (ku *kecUseCase) DeleteKecamatan(id int) (idKecamatan int, err error) {
	idKecamatan, err = ku.kecData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idKecamatan, errors.New("data not found")
		} else {
			return idKecamatan, errors.New("failed to delete data")
		}
	}
	return idKecamatan, nil
}
