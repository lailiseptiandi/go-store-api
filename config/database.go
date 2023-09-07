package config

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	godotenv.Load(path.Join(os.Getenv("HOME"), "../.env"))
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	// dsn := "root:123qwe@tcp(127.0.0.1:3306)/golang-web-toko-online?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}
	return db
}
