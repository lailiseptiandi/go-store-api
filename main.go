package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/config"
	"github.com/lailiseptiandi/golang-toko-online/routers"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.GetConnection()
)

func main() {
	config.LoadEnv()
	config.MigrateDatabase(db)
	config.Seeder(db)
	defer config.DisconnectDB(db)

	r := gin.Default()
	routers.InitRouter(r)
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}
