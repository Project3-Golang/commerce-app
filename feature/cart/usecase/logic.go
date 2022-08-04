package usecase

import (
	"commerce-app/domain"
	"errors"
	"fmt"

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

func (cc *cartUseCase) AddCart(IDUser int, newCart domain.Cart) (domain.Cart, error) {
	if IDUser == -1 {
		return domain.Cart{}, errors.New("invalid user")
	}

	newCart.UserID = IDUser
	fmt.Println("cart", newCart)
	res := cc.cartData.Insert(newCart)

	if res.ID == 0 {
		return domain.Cart{}, errors.New("error insert")
	}
	return res, nil
}

// func (cc *cartUseCase) UpCart(IDCart int, updateData domain.Cart) (domain.Cart, error) {
// 	if IDCart == -1 {
// 		return domain.Cart{}, errors.New("invalid cart")
// 	}

// 	// updateData.UserID = IDNews
// 	res := cc.cartData.Update(IDCart, updateData)
// 	if res.ID == 0 {
// 		return domain.Cart{}, errors.New("error update cart")
// 	}

// 	return res, nil
// }

func (cc *cartUseCase) UpCart(IDCart int, updateData domain.Cart) (domain.Cart, error) {
	if IDCart == -1 {
		return domain.Cart{}, errors.New("invalid cart")
	}

	res := cc.cartData.Update(IDCart, updateData)
	if res.ID == 0 {
		return domain.Cart{}, errors.New("error update")
	}

	return res, nil
}

func (cc *cartUseCase) DelCart(IDCart int) (bool, error) {
	res := cc.cartData.Delete(IDCart)

	if !res {
		return false, errors.New("failed delete")
	}

	return true, nil
}

func (cc *cartUseCase) GetAllC() ([]domain.Cart, error) {
	res := cc.cartData.GetAll()

	if len(res) == 0 {
		return nil, errors.New("no data found")
	}

	return res, nil
}

func (cc *cartUseCase) GetSpecificCart(cartID int) ([]domain.Cart, error) {
	res := cc.cartData.GetCartID(cartID)
	if cartID == -1 {
		return nil, errors.New("error update cart")
	}

	return res, nil
}
