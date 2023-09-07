package config

import (
	"github.com/lailiseptiandi/golang-toko-online/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	createUserAdminSeeder(db)
	createPermissionSeeder(db)
	rolePermissionSeeder(db)
}

func createUserAdminSeeder(db *gorm.DB) {
	password, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	user := models.User{
		ID:       1,
		Name:     "Admin",
		Email:    "admin@admin.com",
		Password: string(password),
		Roles:    "admin",
	}
	_ = db.Debug().FirstOrCreate(&user).Error
}

func createPermissionSeeder(db *gorm.DB) error {

	permissions := []models.Permission{
		models.Permission{
			Name: "category.create",
		},
		models.Permission{
			Name: "category.update",
		},
		models.Permission{
			Name: "category.delete",
		},
		models.Permission{
			Name: "product.create",
		},
		models.Permission{
			Name: "product.update",
		},
		models.Permission{
			Name: "product.delete",
		},
		models.Permission{
			Name: "product_image.create",
		},
		models.Permission{
			Name: "product_image.update",
		},
		models.Permission{
			Name: "product_image.delete",
		},
	}
	tx := db.Begin()

	// Loop through permissions and insert them
	for _, permission := range permissions {
		if err := tx.FirstOrCreate(&permission, models.Permission{Name: permission.Name}).Error; err != nil {
			// Rollback the transaction and return the error
			tx.Rollback()
			return err
		}
	}

	// Commit the transaction
	return tx.Commit().Error
}

func rolePermissionSeeder(db *gorm.DB) error {
	var permissions []models.Permission
	err := db.Model(&models.Permission{}).Find(&permissions).Error
	if err != nil {
		return err

	}
	tx := db.Debug()
	for _, permission := range permissions {
		rolePermsision := models.RolePermission{
			Role:         "admin",
			PermissionID: permission.ID,
		}
		if err := tx.Model(&models.RolePermission{}).FirstOrCreate(&rolePermsision, models.RolePermission{Role: "admin", PermissionID: permission.ID}).Error; err != nil {
			// Rollback the transaction and return the error
			tx.Rollback()
			return err
		}
	}
	return nil
}
