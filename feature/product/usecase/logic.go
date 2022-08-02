package usecase

import (
	"commerce-app/domain"
	"commerce-app/feature/product/data"
	"errors"
	"log"

	validator "github.com/go-playground/validator/v10"
)

type productUseCase struct {
	productData domain.ProductData
	validate    *validator.Validate
}

func New(ud domain.ProductData, v *validator.Validate) domain.ProductUseCase {
	return &productUseCase{
		productData: ud,
		validate:    v,
	}
}
func (pd *productUseCase) GetAllP() ([]domain.Product, error) {
	res := pd.productData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (pd *productUseCase) AddProduct(newProduct domain.Product) (domain.Product, error) {
	var cnv = data.ToLocal(newProduct)
	err := pd.validate.Struct(cnv)
	if err != nil {
		log.Println("Validation error : ", err.Error())
		return domain.Product{}, err
	}

	if err != nil {
		log.Println("error from usecase", err.Error())
		return domain.Product{}, err
	}
	if newProduct.ID == 0 {
		return domain.Product{}, errors.New("cannot insert data")
	}
	return domain.Product, err
}

func (pd *productUseCase) UpProduct(IDProduct int, updateData domain.Product) (domain.Product, error) {

	if IDProduct == -1 {
		return domain.Product{}, errors.New("invalid product")
	}
	result := pd.productData.Update(IDProduct, updateData)

	if result.ID == 0 {
		return domain.Product{}, errors.New("error update product")
	}
	return result, nil
}

func (pd *productUseCase) DelProduct(IDProduct int) (bool, error) {
	res := pd.productData.Delete(IDProduct)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}
