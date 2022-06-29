package usecase

import (
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/repository"
	"errors"
)

type AuthenticationUseCase interface {
	Login(username string, password string) (model.Customer, error)
}

type authenticationUseCase struct {
	customerRepo       repository.CustomerRepository
	userCredentialRepo repository.UserCredentialRepository
}

func (a *authenticationUseCase) Login(username string, password string) (model.Customer, error) {
	userCredential, _ := a.userCredentialRepo.FindByUsernamePassword(username, password)
	customer, err := a.customerRepo.FindFirstWithPreload(map[string]interface{}{"user_credential_id": userCredential.ID}, "UserCredential")
	if err != nil {
		return customer, errors.New("user credential not found")
	}
	return customer, nil
}

func NewAuthenticationUseCase(customerRepo repository.CustomerRepository, userCredentialRepo repository.UserCredentialRepository) AuthenticationUseCase {
	uc := new(authenticationUseCase)
	uc.customerRepo = customerRepo
	uc.userCredentialRepo = userCredentialRepo
	return uc
}
