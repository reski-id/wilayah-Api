package usecase

import (
	"testing"
	"wilayah/domain"
	"wilayah/domain/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{
		UserName: "antonio",
		Password: "antonio123",
		FullName: "Antonio Banderas",
	}

	mockDataEmpty := domain.User{
		ID:       0,
		UserName: "",
		Password: "",
		FullName: "",
	}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$zrIKTTNydPpnWWtn3s6dmeduXInKqj7cVMo6kSTv3cwWBUnPDe/7S"

	invalidData := mockDataEmpty

	noData := mockData
	noData.ID = 0

	t.Run("Sukses Insert Data", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, _ := useCase.AddUser(mockData)
		assert.Equal(t, returnData.ID, res.ID, "Equal")
		assert.NotEqualValues(t, returnData.ID, 0, "not Equal 0")
		assert.NotNil(t, returnData.ID, "not nill")
		repo.AssertExpectations(t)
	})

	t.Run("Cek Validation", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, _ := useCase.AddUser(invalidData)
		assert.NotEqual(t, returnData.ID, res.ID)
		assert.NotNil(t, returnData.UserName, "username not empty")
		assert.NotNil(t, returnData.Password, "password not empty")
		repo.AssertExpectations(t)
	})

	t.Run("Generate Hash Pasword", func(t *testing.T) {
		useCase := New(repo, validator.New())
		res, _ := useCase.AddUser(invalidData)
		assert.Equal(t, returnData.Password, "$2a$10$zrIKTTNydPpnWWtn3s6dmeduXInKqj7cVMo6kSTv3cwWBUnPDe/7S", "pasword Equal")
		assert.NotNil(t, res.Password, "not nill")
		repo.AssertExpectations(t)
	})

}

func TestLoginUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{ID: 1, UserName: "antonio", Password: "123"}
	returnData := domain.User{ID: 1}

	t.Run("Success Login", func(t *testing.T) {
		repo.On("Login", mock.Anything).Return(200, returnData, nil).Once()
		userUseCase := New(repo, validator.New())
		_, err, _ := userUseCase.LoginUser(mockData)

		assert.NotNil(t, err)
		repo.AssertExpectations(t)
	})
}

func TestUpdateUser(t *testing.T) {
	repo := new(mocks.UserData)

	mockData := domain.User{ID: 1, UserName: "antonio", Password: "antonio123", FullName: "Antonio Banderas"}

	returnData := domain.User{ID: 1, UserName: "antonio", Password: "antonio123", FullName: "Antonio Banderas"}

	t.Run("Success Update", func(t *testing.T) {
		repo.On("Update", mock.Anything, mock.Anything).Return(returnData).Once()
		useCase := New(repo, validator.New())
		res, _ := useCase.UpdateUser(1, mockData)
		assert.Equal(t, returnData, res)
		repo.AssertExpectations(t)
	})
}

func TestDeleteUser(t *testing.T) {
	repo := new(mocks.UserData)

	t.Run("Success Delete", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(200, nil).Once()
		useCase := New(repo, validator.New())
		delete, _ := useCase.DeleteUser(1)

		assert.Equal(t, 200, delete)
		repo.AssertExpectations(t)
	})
	t.Run("Data Not Found", func(t *testing.T) {
		repo.On("Delete", mock.Anything).Return(404, nil).Once()
		useCase := New(repo, validator.New())
		delete, _ := useCase.DeleteUser(0)

		assert.Equal(t, 404, delete)
		repo.AssertExpectations(t)
	})
}
