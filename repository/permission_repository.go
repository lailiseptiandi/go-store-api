package repository

import "github.com/lailiseptiandi/golang-toko-online/models"

func (r *repository) ListPermission() ([]models.Permission, error) {
	var permissions []models.Permission
	err := r.db.Model(&models.Permission{}).Find(&permissions).Error
	if err != nil {
		return permissions, err
	}
	return permissions, nil
}
