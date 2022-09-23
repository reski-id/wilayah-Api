package usecase

import (
	"errors"
	"wilayah/domain"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type memUseCase struct {
	memData  domain.MemberData
	validate *validator.Validate
}

func New(kd domain.MemberData, v *validator.Validate) domain.MemberUsecase {
	return &memUseCase{
		memData:  kd,
		validate: v,
	}
}

func (mu *memUseCase) AddMember(newData domain.Member) (idMember int, err error) {
	if newData.NIK == 0 {
		return -1, errors.New("invalid")
	}
	inserted, err := mu.memData.InsertMember(newData)
	return inserted, err
}

func (mu *memUseCase) GetAllMember() ([]domain.Member, error) {
	data, err := mu.memData.GetAll()
	return data, err
}

func (mu *memUseCase) UpdateMember(id int, UpdateData domain.Member) (idMember int, err error) {

	if UpdateData.NIK == 0 {
		return -1, errors.New("invalid")
	}

	data, err := mu.memData.InsertMember(UpdateData)
	return data, err
}

func (mu *memUseCase) DeleteMember(id int) (idMember int, err error) {
	idMember, err = mu.memData.Delete(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return idMember, errors.New("data not found")
		} else {
			return idMember, errors.New("failed to delete data")
		}
	}
	return idMember, nil
}
