package service

import (
	"github.com/lailiseptiandi/golang-toko-online/entity"
	"github.com/lailiseptiandi/golang-toko-online/model"
	"github.com/lailiseptiandi/golang-toko-online/repository"
)

type Service interface {

	// user
	RegisterUser(input entity.RegisterUserInput) (model.User, error)
	LoginUser(input entity.LoginInput) (model.User, error)
	CheckEmailUser(input entity.CheckEmailUser) (bool, error)
	GetUserByID(ID int) (model.User, error)
	UpdateUser(inputID entity.GetUserID, input entity.RegisterUserInput) (model.User, error)

	// category
	GetCategory() ([]model.Category, error)
	CreateCategory(input entity.CreateCategoryInput) (model.Category, error)
	CategoryByID(ID int) (model.Category, error)
	UpdateCategory(inputID entity.CategoryDetailInput, input entity.CreateCategoryInput) (model.Category, error)
	DeleteCategory(ID int) (model.Category, error)

	// product
	GetProduct() ([]model.Product, error)
	CreateProduct(product entity.CreateProduct) (model.Product, error)
	ProductByID(ID int) (model.Product, error)
	UpdateProduct(ID int, input entity.CreateProduct) (model.Product, error)
	DeleteProduct(ID int) error
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
