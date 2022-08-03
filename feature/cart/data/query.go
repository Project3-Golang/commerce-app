package data

import (
	"commerce-app/domain"
	"log"

	"gorm.io/gorm"
)

type cartData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.CartData {
	return &cartData{
		db: db,
	}
}

func (cd *cartData) GetAll() []domain.Cart {
	var data []Cart
	err := cd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

// func (cd *cartData) Insert(newCart domain.Cart) domain.Cart {
// 	cnv := ToLocal(newCart)
// 	err := cd.db.Create(&cnv)
// 	fmt.Println("error", err.Error)
// 	if err.Error != nil {
// 		return domain.Cart{}
// 	}
// 	return cnv.ToDomain()
// }

// func (pd *productData) Update(productID int, updatedProduct domain.Product) domain.Product {
// 	cnv := ToLocal(updatedProduct)
// 	err := pd.db.Model(cnv).Where("ID = ?", productID).Updates(updatedProduct)
// 	if err.Error != nil {
// 		log.Println("Cannot update data", err.Error.Error())
// 		return domain.Product{}
// 	}
// 	cnv.ID = uint(productID)
// 	return cnv.ToDomain()
// }

// func (pd *productData) Delete(productID int) bool {
// 	err := pd.db.Where("ID = ?", productID).Delete(&Product{})
// 	if err.Error != nil {
// 		log.Println("Cannot delete data", err.Error.Error())
// 		return false
// 	}
// 	if err.RowsAffected < 1 {
// 		log.Println("No data deleted", err.Error.Error())
// 		return false
// 	}
// 	return true
// }

// func (pd *productData) GetProductID(productID int) []domain.Product {
// 	var data []Product
// 	err := pd.db.Where("ID = ?", productID).First(&data)

// 	if err.Error != nil {
// 		log.Println("problem data", err.Error.Error())
// 		return nil
// 	}
// 	return ParseToArr(data)
// }
