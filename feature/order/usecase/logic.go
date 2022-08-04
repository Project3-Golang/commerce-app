package usecase

import (
	"commerce-app/domain"
	"errors"

	validator "github.com/go-playground/validator/v10"
)

type orderUseCase struct {
	orderData domain.OrderData
	validate  *validator.Validate
}

func New(ud domain.OrderData, v *validator.Validate) domain.OrderUseCase {
	return &orderUseCase{
		orderData: ud,
		validate:  v,
	}
}

func (pu *orderUseCase) AddOrder(newOrder domain.Order) (domain.Order, error) {

	res := pu.orderData.Insert(newOrder)

	if res.ID == 0 {
		return domain.Order{}, errors.New("error insert data")
	}
	return res, nil
}

// func (pd *productUseCase) GetSpecificProduct(productID int) ([]domain.Product, error) {
// 	res := pd.productData.GetProductID(productID)
// 	if productID == -1 {
// 		return nil, errors.New("error get Product")
// 	}

// 	return res, nil
// }

// func (pd *productUseCase) GetAllP() ([]domain.Product, error) {
// 	res := pd.productData.GetAll()

// 	if len(res) == 0 {
// 		return nil, errors.New("no data found")
// 	}

// 	return res, nil
// }

// func (pd *productUseCase) UpProduct(IDProduct int, updateData domain.Product) (domain.Product, error) {

// 	if IDProduct == -1 {
// 		return domain.Product{}, errors.New("invalid product")
// 	}
// 	result := pd.productData.Update(IDProduct, updateData)

// 	if result.ID == 0 {
// 		return domain.Product{}, errors.New("error update product")
// 	}
// 	return result, nil
// }

// func (pd *productUseCase) DelProduct(IDProduct int) (bool, error) {
// 	res := pd.productData.Delete(IDProduct)

// 	if !res {
// 		return false, errors.New("failed delete")
// 	}

// 	return true, nil
// }
