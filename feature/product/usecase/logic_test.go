package usecase

import (
	"commerce-app/domain"
	"commerce-app/domain/mocks"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddProduct(t *testing.T) {
	repo := new(mocks.ProductData)

	mockData := domain.Product{
		ID:    1,
		Name:  "Product Baru",
		Price: 1000,
		Stock: 10,
	}

	// mockDataEmpty := domain.Product{
	// 	ID:    0,
	// 	Name:  "",
	// 	Price: 0,
	// 	Stock: 0,
	// }

	returnData := mockData

	// invalidData := mockDataEmpty

	noData := mockData
	noData.ID = 0

	t.Run("Sukses Insert Product", func(t *testing.T) {
		repo.On("Insert", mock.Anything).Return(returnData, nil).Once()
		useCase := New(repo, validator.New())
		res, _ := useCase.AddProduct(mockData)

		assert.Equal(t, returnData.ID, res.ID, "Equal")
		assert.NotEqualValues(t, returnData.ID, 0, "not Equal 0")
		assert.NotNil(t, returnData.ID, "not nill")

		repo.AssertExpectations(t)
	})

}
