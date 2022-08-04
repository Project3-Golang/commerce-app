package data

import (
	"commerce-app/domain"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type orderData struct {
	db *gorm.DB
}

func New(db *gorm.DB) domain.OrderData {
	return &orderData{
		db: db,
	}
}

func (od *orderData) Insert(newOrder domain.Order) domain.Order {
	cnv := ToLocal(newOrder)
	err := od.db.Create(&cnv)
	fmt.Println("error", err.Error)
	if err.Error != nil {
		return domain.Order{}
	}
	return cnv.ToDomain()
}
func (od *orderData) Update(orderID int, updatedOrder domain.Order) domain.Order {
	cnv := ToLocal(updatedOrder)
	err := od.db.Model(cnv).Where("ID = ?", orderID).Updates(updatedOrder)
	if err.Error != nil {
		log.Println("Cannot update data", err.Error.Error())
		return domain.Order{}
	}
	cnv.ID = uint(orderID)
	return cnv.ToDomain()
}

// func (od *orderData) Update(orderID int, updatedOrder domain.Order) domain.Order {
// 	cnv := ToLocal(updatedOrder)
// 	err := od.db.Model(cnv).Where("ID = ?", orderID).Updates(updatedOrder)
// 	if err.Error != nil {
// 		log.Println("Cannot update data", err.Error.Error())
// 		return domain.Order{}
// 	}
// 	cnv.ID = uint(orderID)
// 	return cnv.ToDomain()
// }

func (od *orderData) Delete(orderID int) bool {
	err := od.db.Where("ID = ?", orderID).Delete(&Order{})
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

func (od *orderData) GetAll() []domain.Order {
	var data []Order
	err := od.db.Find(&data)

	if err.Error != nil {
		log.Println("error on select data", err.Error.Error())
		return nil
	}

	return ParseToArr(data)
}

func (od *orderData) GetOrderID(orderID int) []domain.Order {
	var data []Order
	err := od.db.Where("ID = ?", orderID).First(&data)

	if err.Error != nil {
		log.Println("problem data", err.Error.Error())
		return nil
	}
	return ParseToArr(data)
}
