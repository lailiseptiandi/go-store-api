package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lailiseptiandi/golang-toko-online/config"
	"github.com/lailiseptiandi/golang-toko-online/handler"
	"github.com/lailiseptiandi/golang-toko-online/middleware"
	"github.com/lailiseptiandi/golang-toko-online/repository"
	"github.com/lailiseptiandi/golang-toko-online/service"
)

func main() {
	globalRepository := repository.NewRepository()

	globalService := service.NewService(globalRepository)
	authService := config.NewService()

	globalHandler := handler.NewHandler(globalService, authService)
	// categoryHandler := handler.NewHandlerCategory(globalService, authService)

	router := gin.Default()

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

	router.Run()
}
