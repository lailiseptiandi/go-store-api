package repository

import "github.com/lailiseptiandi/golang-toko-online/model"

func (r *repository) GetCategory() ([]model.Category, error) {

	var categories []model.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err

	}
	return categories, nil
}

func (r *repository) CreateCategory(category model.Category) (model.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err

	}
	return category, nil
}

func (r *repository) CategoryBYID(ID int) (model.Category, error) {
	var category model.Category
	err := r.db.Where("id=?", ID).Find(&category).Error
	if err != nil {
		return category, err

	}
	return category, nil
}

func (r *repository) UpdateCategory(category model.Category) (model.Category, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategory(ID int) (model.Category, error) {
	var category model.Category
	err := r.db.Where("id = ?", ID).Find(&category).Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
