package usecase

import (
	"commerce-app/domain"
	"errors"

	validator "github.com/go-playground/validator/v10"
)

type cartUseCase struct {
	cartData domain.CartData
	validate *validator.Validate
}

func New(ud domain.CartData, v *validator.Validate) domain.CartUseCase {
	return &cartUseCase{
		cartData: ud,
		validate: v,
	}
}
func (pu *cartUseCase) GetAllC() ([]domain.Cart, error) {
	res := pu.cartData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

// func (pd *productUseCase) AddProduct(newProduct domain.Product) (domain.Product, error) {

// 	res := pd.productData.Insert(newProduct)

// 	if res.ID == 0 {
// 		return domain.Product{}, errors.New("error insert data")
// 	}
// 	return res, nil
// }

// func (pd *productUseCase) GetSpecificProduct(productID int) ([]domain.Product, error) {
// 	res := pd.productData.GetProductID(productID)
// 	if productID == -1 {
// 		return nil, errors.New("error get Product")
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
