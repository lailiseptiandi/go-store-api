package service

import (
	"github.com/lailiseptiandi/go-store-api/entity"
	"github.com/lailiseptiandi/go-store-api/models"
	"github.com/lailiseptiandi/go-store-api/repository"
)

type Service interface {

	// user
	RegisterUser(input entity.RegisterUserInput) (models.User, error)
	LoginUser(input entity.LoginInput) (models.User, error)
	CheckEmailUser(input entity.CheckEmailUser) (bool, error)
	GetUserByID(ID int) (models.User, error)
	UpdateUser(inputID entity.GetUserID, input entity.RegisterUserInput) (models.User, error)

	// category
	GetCategory() ([]models.Category, error)
	CreateCategory(input entity.CreateCategoryInput) (models.Category, error)
	CategoryByID(ID int) (models.Category, error)
	UpdateCategory(inputID entity.CategoryDetailInput, input entity.CreateCategoryInput) (models.Category, error)
	DeleteCategory(ID int) (models.Category, error)

	// product
	GetProduct() ([]models.Product, error)
	CreateProduct(product entity.CreateProduct) (models.Product, error)
	ProductByID(ID int) (models.Product, error)
	UpdateProduct(ID int, input entity.CreateProduct) (models.Product, error)
	DeleteProduct(ID int) error

	// product image
	GetImageByProduct(int) ([]models.ProductImage, error)
	StoreImageByProduct(models.ProductImage) (models.ProductImage, error)
}

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}
