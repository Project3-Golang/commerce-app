package data

import (
	"commerce-app/domain"
	"fmt"
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

func (cd *cartData) Insert(newCart domain.Cart) domain.Cart {
	cnv := ToLocal(newCart)
	err := cd.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Cart{}
	}
	return cnv.ToDomain()
}

// func (cd *cartData) Update(cartID int, updatedCart domain.Cart) domain.Cart {
// 	cnv := ToLocal(updatedCart)
// 	err := cd.db.Model(cnv).Where("ID = ?", cartID).Updates(updatedCart)
// 	if err.Error != nil {
// 		log.Println("Cannot update data", err.Error.Error())
// 		return domain.Cart{}
// 	}
// 	cnv.ID = uint(cartID)
// 	return cnv.ToDomain()
// }

func (cd *cartData) Update(cartID int, updatedCart domain.Cart) domain.Cart {
	cnv := ToLocal(updatedCart)
	err := cd.db.Model(cnv).Where("ID = ?", cartID).Updates(updatedCart)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Cart{}
	}
	cnv.ID = uint(cartID)
	return cnv.ToDomain()
}

func (cd *cartData) Delete(cartID int) bool {
	err := cd.db.Where("ID = ?", cartID).Delete(&Cart{})
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

func (cd *cartData) GetAll() []domain.Cart {
	var data []Cart
	err := cd.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (cd *cartData) GetCartID(cartID int) []domain.Cart {
	var data []Cart
	err := cd.db.Where("ID = ?", cartID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
