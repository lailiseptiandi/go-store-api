package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/routers"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	db          *gorm.DB      = config.MysqlConnection()
	redisClient *redis.Client = config.RedisConnection()
)

func main() {
	config.LoadEnv()
	config.MigrateDatabase(db)
	config.Seeder(db)
	defer config.DisconnectDB(db)

	r := gin.Default()
	r.Use(corsMiddleware())
	routers.InitRouter(r, db, redisClient)
	r.Run(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, refresh_token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
