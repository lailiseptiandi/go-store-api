package repository

import "github.com/lailiseptiandi/go-store-api/models"

func (r *repository) SaveUser(user models.User) (models.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (models.User, error) {

	var user models.User
	err := r.db.Where("email= ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByID(ID int) (models.User, error) {

	var user models.User
	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UpdateUser(user models.User) (models.User, error) {
	err := r.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil

}

func (r *repository) ListUser() ([]models.User, error) {
	var user []models.User

	err := r.db.Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
