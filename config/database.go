package config

import (
	"fmt"
	"os"

	"github.com/lailiseptiandi/golang-toko-online/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	// godotenv.Load(path.Join(os.Getenv("HOME"), "../.env"))
	LoadEnv()

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_DATABASE")
	// dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, db_name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}
	return db
}

func MigrateDatabase(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.ProductImage{})
}

func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
