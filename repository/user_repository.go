package repository

import (
	"github.com/lailiseptiandi/golang-toko-online/model"
)

func (r *repository) SaveUser(user model.User) (model.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (model.User, error) {

	var user model.User
	err := r.db.Where("email= ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (model.User, error) {
	var user model.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err

	}

	return user, nil
}

func (r *repository) UpdateUser(user model.User) (model.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}
