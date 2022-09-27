package repository

import (
	"github.com/lailiseptiandi/golang-toko-online/config"
	"github.com/lailiseptiandi/golang-toko-online/model"
	"gorm.io/gorm"
)

type (
	Repository interface {

		// User
		SaveUser(user model.User) (model.User, error)
		FindByEmail(email string) (model.User, error)
		FindByID(ID int) (model.User, error)
		UpdateUser(user model.User) (model.User, error)

		// category
		GetCategory() ([]model.Category, error)
		CreateCategory(category model.Category) (model.Category, error)
		CategoryBYID(ID int) (model.Category, error)
		UpdateCategory(category model.Category) (model.Category, error)
		DeleteCategory(ID int) (model.Category, error)

		// product
	}
)
type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	return &repository{config.GetConnection()}
}
