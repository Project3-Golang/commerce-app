package usecase

import (
	"commerce-app/domain"
	"errors"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type productUseCase struct {
	productData domain.ProductData
}

func New(ud domain.ProductData) domain.ProductUseCase {
	return &productUseCase{
		productData: ud,
	}
}

func (pu *productUseCase) AddProduct(newProduct domain.Product) (domain.Product, error) {

	fmt.Println("news", newProduct)
	res, _ := pu.productData.Insert(newProduct)

	if res.ID == 0 {
		return domain.Product{}, errors.New("error insert product")
	}
	return res, nil
}

func (pu *productUseCase) GetProduct(id int) (domain.Product, error) {
	res, _ := pu.productData.GetSpecific(id)
	if id == -1 {
		return domain.Product{}, errors.New("error update news")
	}

	return res, nil
}

func (pu *productUseCase) DeleteProduct(id int) (row int, err error) {
	row, err = pu.productData.Delete(id)
	if err != nil {
		log.Println("delete usecase error", err.Error())
		if err == gorm.ErrRecordNotFound {
			return row, errors.New("data not found")
		} else {
			return row, errors.New("server error")
		}
	}
	return row, nil
}

func (pu *productUseCase) GetAll() ([]domain.Product, error) {
	res, _ := pu.productData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (pu *productUseCase) UpdateProduct(IDProduct int, updateData domain.Product) (domain.Product, error) {
	if IDProduct == -1 {
		return domain.Product{}, errors.New("invalid product")
	}

	res := pu.productData.Update(IDProduct, updateData)
	if res.ID == 0 {
		return domain.Product{}, errors.New("error update product")
	}

	return res, nil
}
