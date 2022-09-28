package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gosimple/slug"
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/model"
)

func (s *service) GetProduct() ([]model.Product, error) {

	listProduct, err := s.repository.GetProduct()
	if err != nil {
		return listProduct, err
	}

	return listProduct, nil
}

func (s *service) CreateProduct(product entity.CreateProduct) (model.Product, error) {

	products := model.Product{
		CategoryID:   product.CategoryID,
		Name:         product.Name,
		Slug:         slug.Make(strings.ToLower(product.Name)),
		Descriptions: product.Descriptions,
		Price:        product.Price,
		Quantity:     product.Qty,
	}
	newProduct, err := s.repository.CreateProduct(products)
	if err != nil {
		return newProduct, err
	}
	return newProduct, nil
}

func (s *service) ProductByID(ID int) (model.Product, error) {
	products, err := s.repository.ProductByID(ID)
	fmt.Println("ini product ", products)
	if err != nil {
		return products, err
	}
	if products.ID == 0 {
		return products, errors.New("product not found")
	}
	return products, nil
}

func (s *service) UpdateProduct(ID int, product entity.CreateProduct) (model.Product, error) {
	checkProduct, err := s.repository.ProductByID(ID)
	if err != nil {
		return checkProduct, err
	}
	products := model.Product{
		CategoryID:   product.CategoryID,
		Name:         product.Name,
		Slug:         slug.Make(strings.ToLower(product.Name)),
		Descriptions: product.Descriptions,
		Price:        product.Price,
		Quantity:     product.Qty,
	}

	updateProduct, err := s.repository.UpdateProduct(checkProduct.ID, products)
	if err != nil {
		return updateProduct, err
	}
	return updateProduct, nil
}

func (s *service) DeleteProduct(ID int) error {
	err := s.repository.DeleteProduct(ID)
	if err != nil {
		return err
	}
	return nil
}
