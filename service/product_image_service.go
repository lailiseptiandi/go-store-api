package service

import "github.com/lailiseptiandi/golang-toko-online/models"

func (s *service) GetImageByProduct(ID int) ([]models.ProductImage, error) {
	image, err := s.repository.ProductImageByID(ID)
	if err != nil {
		return image, err
	}
	return image, nil
}

func (s *service) StoreImageByProduct(productImage models.ProductImage) (models.ProductImage, error) {
	image, err := s.repository.StoreProductImage(productImage)
	if err != nil {
		return image, err
	}
	return image, nil

}
