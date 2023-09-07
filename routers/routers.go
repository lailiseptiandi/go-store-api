package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/handler"
	"github.com/lailiseptiandi/go-store-api/middleware"
	"github.com/lailiseptiandi/go-store-api/repository"
	"github.com/lailiseptiandi/go-store-api/service"
)

func InitRouter(router *gin.Engine) {
	globalRepository := repository.NewRepository()

	globalService := service.NewService(globalRepository)
	authService := config.NewService()

	globalHandler := handler.NewHandler(globalService, authService)
	// categoryHandler := handler.NewHandlerCategory(globalService, authService)

	router.POST("/register", globalHandler.RegisterUser)
	router.POST("/login", globalHandler.LoginUser)
	router.POST("/check-email", globalHandler.CheckEmailAvailable)
	router.PUT("/users/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.UpdateUser)
	router.GET("/users/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.GetUserByID)

	// category
	router.GET("category", globalHandler.GetCategory)
	router.POST("category", middleware.AuthMiddleware(authService, globalService), globalHandler.CreateCategory)
	router.GET("category/:id", globalHandler.DetailCategory)
	router.PUT("category/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.UpdateCategory)
	router.DELETE("category/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.DeleteGategory)

	// product
	router.GET("product", globalHandler.GetProduct)
	router.POST("product", middleware.AuthMiddleware(authService, globalService), globalHandler.CreateProduct)
	router.GET("product/:id", globalHandler.ProductByID)
	router.PUT("product/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.UpdateProduct)
	router.DELETE("product/:id", middleware.AuthMiddleware(authService, globalService), globalHandler.DeleteProduct)

	// product image
	router.GET("image/:id", globalHandler.GetImageByProduct)
}
