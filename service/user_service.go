package service

import (
	"errors"

	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/models"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) RegisterUser(input entity.RegisterUserInput) (models.User, error) {
	user := models.User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Address = input.Address
	user.Roles = "user"
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(password)

	newUser, err := s.repository.SaveUser(user)
	if err != nil {
		return newUser, err

	}

	return newUser, nil

}

func (s *service) LoginUser(input entity.LoginInput) (models.User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found that email")
	}

	// membandingkan database input
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *service) CheckEmailUser(input entity.CheckEmailUser) (bool, error) {

	var email = input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *service) GetUserByID(ID int) (models.User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found ID")
	}
	return user, nil
}
func (s *service) UpdateUser(inputID entity.GetUserID, input entity.RegisterUserInput) (models.User, error) {

	user, err := s.repository.FindByID(inputID.ID)
	if user.ID != inputID.ID {
		return user, err
	}
	user.Name = input.Name
	user.Email = input.Email
	user.Roles = "user"
	user.Address = input.Address
	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Token = "123awdasdsad"
	user.Password = string(password)
	updateUser, err := s.repository.UpdateUser(user)

	if err != nil {
		return updateUser, err
	}

	return updateUser, nil

}
