package usecase

import (
	"commerce-app/domain"
	"commerce-app/feature/product/data"
	"errors"
	"log"

	validator "github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
	if inserted.ID == 0 {
		return domain.Product{}, errors.New("cannot insert data")
	}
	return inserted, nil
}
func (pd *productUseCase) UpdateProduct(id int, updateProfile domain.User) (domain.User, error) {

	if id == -1 {
		return domain.User{}, errors.New("invalid user")
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte(updateProfile.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("error encrypt password", err)
		return domain.User{}, err
	}
	updateProfile.Password = string(hashed)
	result := ud.userData.Update(id, updateProfile)

	if result.ID == 0 {
		return domain.User{}, errors.New("error update user")
	}
	return result, nil
}

func (pd *productUseCase) LoginUser(userLogin domain.User) (response int, data domain.User, err error) {
	response, data, err = ud.userData.Login(userLogin)

	return response, data, err
}

func (pd *productUseCase) GetProfile(id int) (domain.User, error) {
	data, err := ud.userData.GetSpecific(id)

	if err != nil {
		log.Println("Use case", err.Error())
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, errors.New("data not found")
		} else {
			return domain.User{}, errors.New("server error")
		}
	}

	return data, nil
}

func (pd *productUseCase) DeleteUser(id int) (row int, err error) {
	row, err = ud.userData.Delete(id)
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
