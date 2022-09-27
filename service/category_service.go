package service

import (
	"errors"
	"strings"

	"github.com/gosimple/slug"
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/model"
)

func (s *service) GetCategory() ([]model.Category, error) {

	listCategory, err := s.repository.GetCategory()
	if err != nil {
		return listCategory, err
	}

	return listCategory, nil
}

func (s *service) CreateCategory(input entity.CreateCategoryInput) (model.Category, error) {

	category := model.Category{}
	category.Name = input.Name
	category.Slug = slug.Make(strings.ToLower(input.Name))

	newCategory, err := s.repository.CreateCategory(category)
	if err != nil {
		return newCategory, err
	}

	return newCategory, nil
}

func (s *service) CategoryByID(ID int) (model.Category, error) {
	category, err := s.repository.CategoryBYID(ID)
	if err != nil {
		return category, err
	}

	if category.ID == 0 {
		return category, errors.New("category not found")
	}
	return category, nil
}

func (s *service) UpdateCategory(inputID entity.CategoryDetailInput, input entity.CreateCategoryInput) (model.Category, error) {

	category, err := s.repository.CategoryBYID(inputID.ID)
	if category.ID != inputID.ID {
		return category, err
	}
	category.Name = input.Name
	category.Slug = slug.Make(strings.ToLower(input.Name))

	updateCategory, err := s.repository.UpdateCategory(category)
	if err != nil {
		return updateCategory, err
	}

	return updateCategory, nil
}

func (s *service) DeleteCategory(ID int) (model.Category, error) {
	category, err := s.repository.DeleteCategory(ID)
	if err != nil {
		return category, err
	}
	if category.ID == 0 {
		return category, errors.New("category not found")
	}
	return category, nil
}
