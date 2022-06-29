package repository

import (
	"enigmacamp.com/golang-gorm/model"
	"enigmacamp.com/golang-gorm/utils"
	"errors"
	"gorm.io/gorm"
)

type UserCredentialRepository interface {
	FindByUsername(username string) (model.UserCredential, error)
	FindByUsernamePassword(username string, password string) (model.UserCredential, error)
}

type userCredentialRepository struct {
	db *gorm.DB
}

func (u *userCredentialRepository) FindByUsername(username string) (model.UserCredential, error) {
	var userCredential model.UserCredential
	result := u.db.Where("username = ?", username).First(&userCredential)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return userCredential, nil
		} else {
			return userCredential, errors.New("username not found")
		}
	}
	return userCredential, nil

}

func (u *userCredentialRepository) FindByUsernamePassword(username string, password string) (model.UserCredential, error) {
	var userCredential model.UserCredential
	user, err := u.FindByUsername(username)
	if err != nil {
		return userCredential, errors.New("username not found")
	}
	pwdCheck := utils.CheckPasswordHash(password, user.Password)
	if pwdCheck {
		return user, nil
	} else {
		return user, errors.New("password don't match")
	}
}

func NewUserCredentialRepository(db *gorm.DB) UserCredentialRepository {
	repo := new(userCredentialRepository)
	repo.db = db
	return repo
}
