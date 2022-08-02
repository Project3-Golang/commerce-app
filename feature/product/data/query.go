package data

import (
	"commerce-app/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type productData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.ProductData {
	return &productData{
		db: db,
	}
}

func (pd *productData) Insert(dataProduct domain.Product) domain.Product {
	cnv := ToLocal(dataProduct)
	err := pd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Product{}
	}
	return cnv.ToDomain()
}

func (pd *productData) Update(productID int, updatedProduct domain.Product) domain.Product {
	cnv := ToLocal(updatedProduct)
	err := pd.db.Model(cnv).Where("ID = ?", productID).Updates(updatedProduct)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Product{}
	}
	cnv.ID = uint(productID)
	return cnv.ToDomain()
}

func (pd *productData) Delete(productID int) bool {
	err := pd.db.Where("ID = ?", productID).Delete(&Product{})
	if err.Error != nil {
		log.Println("Cannot delete data", err.Error.Error())
		return false
	}
	if err.RowsAffected < 1 {
		log.Println("No data deleted", err.Error.Error())
		return false
	}
	return true
}

func (pd *productData) GetAll() []domain.Product {
	var data []Product
	err := pd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (pd *productData) GetProductID(productID int) []domain.Product {
	var data []Product
	err := pd.db.Where("ID = ?", productID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
