package usecase

import (
	"commerce-app/domain"
	"commerce-app/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddUser(t *testing.T) {
	repo := new(mocks.UserData)
	mockData := domain.User{
		UserName: "antonio",
		Email:    "antonio@gmail.com",
		Password: "antonio123",
		Role:     "admin",
		FullName: "Antonio Banderas",
	}

	mockDataEmpty := domain.User{
		ID:       0,
		UserName: "",
		Email:    "",
		Password: "",
		Role:     "",
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
		assert.NotNil(t, returnData.Email, "email not empty")
		assert.NotNil(t, returnData.Password, "password not empty")
		assert.NotNil(t, returnData.FullName, "full name not empty")
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

}

func TestUpdateUser(t *testing.T) {

}
