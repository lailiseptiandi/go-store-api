package repository

import "github.com/lailiseptiandi/golang-toko-online/model"

func (r *repository) GetProduct() ([]model.Product, error) {
	var product []model.Product
	err := r.db.Table("products").Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil
}

func (r *repository) CreateProduct(products model.Product) (model.Product, error) {
	err := r.db.Table("products").Create(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) ProductByID(ID int) (model.Product, error) {
	var product model.Product
	err := r.db.Where("id = ?", ID).Find(&product).Error
	if err != nil {
		return product, err
	}
	return product, nil

}

func (r *repository) UpdateProduct(ID int, products model.Product) (model.Product, error) {
	err := r.db.Table("products").Where("id= ?", ID).Updates(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (r *repository) DeleteProduct(ID int) error {
	var product model.Product
	err := r.db.Table("products").Where("id = ?", ID).Delete(&product).Error
	if err != nil {
		return err
	}
	return nil

}
