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

func (pu *orderUseCase) AddOrder(IDUser int, newOrder domain.Order) (domain.Order, error) {
	if IDUser == -1 {
		return domain.Order{}, errors.New("invalid user")
	}

	newOrder.UserID = IDUser
	res := pu.orderData.Insert(newOrder)

	if res.ID == 0 {
		return domain.Order{}, errors.New("error insert")
	}
	return res, nil
}

func (pu *orderUseCase) GetSpecificOrder(orderID int) ([]domain.Order, error) {
	res := pu.orderData.GetOrderID(orderID)
	if orderID == -1 {
		return nil, errors.New("error get Order")
	}

	return res, nil
}

func (pu *orderUseCase) GetAllO() ([]domain.Order, error) {
	res := pu.orderData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (pu *orderUseCase) UpOrder(IDOrder int, updateData domain.Order) (domain.Order, error) {

	if IDOrder == -1 {
		return domain.Order{}, errors.New("invalid order")
	}
	result := pu.orderData.Update(IDOrder, updateData)

	if result.ID == 0 {
		return domain.Order{}, errors.New("error update")
	}
	return result, nil
}

func (pu *orderUseCase) DelOrder(IDOrder int) (bool, error) {
	res := pu.orderData.Delete(IDOrder)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (pu *orderUseCase) GetmyOrder(userID int) ([]domain.Order, error) {
	res := pu.orderData.GetOrderbyuser(userID)
	if userID == -1 {
		return nil, errors.New("error get order")
	}

	return res, nil
}
