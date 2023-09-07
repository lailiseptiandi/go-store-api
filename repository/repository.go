package repository

import (
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/models"
	"gorm.io/gorm"
)

type (
	Repository interface {

		// User
		SaveUser(user models.User) (models.User, error)
		FindByEmail(email string) (models.User, error)
		FindByID(ID int) (models.User, error)
		UpdateUser(user models.User) (models.User, error)

		// repository
		ListPermission() ([]models.Permission, error)

		// category
		GetCategory() ([]models.Category, error)
		CreateCategory(category models.Category) (models.Category, error)
		CategoryBYID(ID int) (models.Category, error)
		UpdateCategory(category models.Category) (models.Category, error)
		DeleteCategory(ID int) (models.Category, error)

		// product
		GetProduct() ([]models.Product, error)
		CreateProduct(product models.Product) (models.Product, error)
		ProductByID(ID int) (models.Product, error)
		UpdateProduct(ID int, product models.Product) (models.Product, error)
		DeleteProduct(ID int) error

		// product images
		GetProductImage() ([]models.ProductImage, error)
		StoreProductImage(models.ProductImage) (models.ProductImage, error)
		ProductImageByID(int) ([]models.ProductImage, error)
		UpdateProductImage(int, models.ProductImage) (models.ProductImage, error)
		DeleteProductImage(int) error
	}
)
type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{config.GetConnection()}
}
