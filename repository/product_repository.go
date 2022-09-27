package repository

import "github.com/lailiseptiandi/golang-toko-online/model"

func (r *repository) GetProduct() ([]model.Product, error) {
	var product []model.Product
	err := r.db.Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}
