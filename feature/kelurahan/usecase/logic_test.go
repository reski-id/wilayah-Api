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

func TestAddKelurahan(t *testing.T) {
	repo := new(mocks.KelurahanData)
	insertData := domain.Kelurahan{
		ID:            1,
		KelurahanName: "Lurah Sepakat",
		KecamatanID:   1,
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertKelurahan", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddKelurahan(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Insert Empty data", func(t *testing.T) {
		repo.On("InsertKelurahan", mock.Anything, mock.Anything).Return(-1, errors.New("invalid")).Once()
		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.KelurahanName = ""

		res, err := useCase.AddKelurahan(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid").Error())
	})
}

func TestDeleteKelurahan(t *testing.T) {
	repo := new(mocks.KelurahanData)
	insertDate := domain.Kelurahan{
		ID:            1,
		KelurahanName: "Lurah Sepakat",
		KecamatanID:   1,
	}

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		usecase := New(repo, validator.New())

		res, err := usecase.DeleteKelurahan(insertDate.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, fmt.Errorf("failed to delete data")).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKelurahan(1)
		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("failed to delete data"))
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, gorm.ErrRecordNotFound).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKelurahan(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllKelurahan(t *testing.T) {
	repo := new(mocks.KelurahanData)

	returnData := []domain.Kelurahan{{
		ID:            1,
		KelurahanName: "Lurah Sepakat",
		KecamatanID:   1,
	}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllKelurahan()

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, nil).Once()
		useCase := New(repo, validator.New())
		result, _ := useCase.GetAllKelurahan()
		assert.Equal(t, []domain.Kelurahan(nil), result)
		repo.AssertExpectations(t)
	})
}
