package usecase

import (
	"errors"
	"fmt"
	"testing"
	"wilayah/domain"
	"wilayah/domain/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestAddKota(t *testing.T) {
	repo := new(mocks.KotaData)
	insertData := domain.Kota{
		ID:         1,
		KotaName:   "Kota Tua",
		ProvinsiID: 1,
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertKota", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddKota(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Insert Empty data", func(t *testing.T) {
		repo.On("InsertKota", mock.Anything, mock.Anything).Return(-1, errors.New("invalid")).Once()
		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.KotaName = ""

		res, err := useCase.AddKota(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid").Error())
		// repo.AssertExpectations(t)
	})
}

func TestDeleteKota(t *testing.T) {
	repo := new(mocks.KotaData)
	insertDate := domain.Kota{
		ID:         1,
		KotaName:   "Kota Tua",
		ProvinsiID: 1,
	}

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		usecase := New(repo, validator.New())

		res, err := usecase.DeleteKota(insertDate.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, fmt.Errorf("failed to delete data")).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKota(1)
		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("failed to delete data"))
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, gorm.ErrRecordNotFound).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKota(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllKota(t *testing.T) {
	repo := new(mocks.KotaData)

	returnData := []domain.Kota{{
		ID:         1,
		KotaName:   "Kota Tua",
		ProvinsiID: 1,
	}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllKota()

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, nil).Once()
		useCase := New(repo, validator.New())
		result, _ := useCase.GetAllKota()
		assert.Equal(t, []domain.Kota(nil), result)
		repo.AssertExpectations(t)
	})
}
