package repository

import "github.com/lailiseptiandi/golang-toko-online/models"

func (r *repository) GetCategory() ([]models.Category, error) {

	var categories []models.Category
	err := r.db.Find(&categories).Error
	if err != nil {
		return categories, err

	}
	return categories, nil
}

func (r *repository) CreateCategory(category models.Category) (models.Category, error) {
	err := r.db.Create(&category).Error
	if err != nil {
		return category, err

	}
	return category, nil
}

func (r *repository) CategoryBYID(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.Where("id=?", ID).Find(&category).Error
	if err != nil {
		return category, err

	}
	return category, nil
}

func (r *repository) UpdateCategory(category models.Category) (models.Category, error) {
	err := r.db.Save(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *repository) DeleteCategory(ID int) (models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ?", ID).Find(&category).Delete(&category).Error

	if err != nil {
		return category, err
	}

	return category, nil
}
