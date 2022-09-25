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

func TestAddKecamatan(t *testing.T) {
	repo := new(mocks.KecamatanData)
	insertData := domain.Kecamatan{
		ID:            1,
		KecamatanName: "Kec. Baru",
		KotaID:        2,
	}

	t.Run("Success Insert", func(t *testing.T) {
		repo.On("InsertKecamatan", mock.Anything).Return(1, nil).Once()

		useCase := New(repo, validator.New())

		res, err := useCase.AddKecamatan(insertData)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Insert Empty data", func(t *testing.T) {
		repo.On("InsertKecamatan", mock.Anything, mock.Anything).Return(-1, errors.New("invalid")).Once()
		useCase := New(repo, validator.New())
		dummy := insertData
		dummy.KecamatanName = ""

		res, err := useCase.AddKecamatan(dummy)
		assert.NotNil(t, err)
		assert.Equal(t, -1, res)
		assert.EqualError(t, err, errors.New("invalid").Error())
	})
}

func TestDeleteKecamatan(t *testing.T) {
	repo := new(mocks.KecamatanData)
	insertDate := domain.Kecamatan{
		ID:            1,
		KecamatanName: "Kec. Baru",
		KotaID:        2,
	}

	t.Run("Success delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(1, nil).Once()

		usecase := New(repo, validator.New())

		res, err := usecase.DeleteKecamatan(insertDate.ID)
		assert.Nil(t, err)
		assert.Equal(t, 1, res)
		repo.AssertExpectations(t)
	})

	t.Run("Failed delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, fmt.Errorf("failed to delete data")).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKecamatan(1)
		assert.NotNil(t, err)
		assert.Equal(t, err, fmt.Errorf("failed to delete data"))
		repo.AssertExpectations(t)
	})

	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(0, gorm.ErrRecordNotFound).Once()

		usecase := New(repo, validator.New())

		_, err := usecase.DeleteKecamatan(1)
		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestGetAllKecamatan(t *testing.T) {
	repo := new(mocks.KecamatanData)

	returnData := []domain.Kecamatan{{
		ID:            1,
		KecamatanName: "Kec. Baru",
		KotaID:        2,
	}}

	t.Run("Success Get All", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, error := useCase.GetAllKecamatan()

		assert.Equal(t, error, nil)
		assert.GreaterOrEqual(t, len(res), 1)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("GetAll", mock.Anything).Return(nil, nil).Once()
		useCase := New(repo, validator.New())
		result, _ := useCase.GetAllKecamatan()
		assert.Equal(t, []domain.Kecamatan(nil), result)
		repo.AssertExpectations(t)
	})
}
