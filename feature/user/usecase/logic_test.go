package usecase

import (
	"backend/domain/mocks"
	"commerce-app/domain"
	"testing"
)

func TestAddUser(t *testing.T) {
	repo = new(mocks.UserData)

	mockData := domain.User{
		ID:       1,
		UserName: "thanos",
		Email:    "lukman@gmail.com",
		Password: "madtitan",
		Role:     "admin",
		FullName: "Thanos",
	}

	mockDataEmpty := domain.User{
		ID:       0,
		UserName: "",
		Email:    "",
		Password: "",
		Role:     "",
		FullName: "",
	}

	returnData := mockData
	returnData.ID = 1
	returnData.Password = "$2a$10$SrMvwwY/QnQ4nZunBvGOuOm2U1w8wcAENOoAMI7l8xH7C1Vmt5mru"

}
