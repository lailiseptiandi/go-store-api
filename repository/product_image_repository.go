package repository

import "github.com/lailiseptiandi/go-store-api/models"

func (r *repository) GetProductImage() ([]models.ProductImage, error) {
	var productImage []models.ProductImage
	err := r.db.Table("product_images").Find(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
}

func (r *repository) StoreProductImage(productImage models.ProductImage) (models.ProductImage, error) {
	err := r.db.Table("product_images").Create(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
}

func (r *repository) ProductImageByID(ID int) ([]models.ProductImage, error) {
	var productImage []models.ProductImage
	err := r.db.Table("product_images").Where("id =?", ID).Find(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
}

func (r *repository) UpdateProductImage(ID int, productImage models.ProductImage) (models.ProductImage, error) {
	err := r.db.Table("product_images").Where("id =?", ID).Updates(&productImage).Error
	if err != nil {
		return productImage, err
	}
	return productImage, nil
}

func (r *repository) DeleteProductImage(ID int) error {
	var productImage models.ProductImage
	err := r.db.Table("product_images").Where("id =?", ID).Delete(&productImage).Error
	if err != nil {
		return err
	}
	return nil
}
