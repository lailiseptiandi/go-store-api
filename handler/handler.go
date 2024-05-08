package handler

import (
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/service"
	"github.com/redis/go-redis/v9"
)

type globalHandler struct {
	globalService service.Service
	authService   config.Service
	redisClient   *redis.Client
}

func NewHandler(globalService service.Service, authService config.Service, redisClient *redis.Client) *globalHandler {
	return &globalHandler{globalService, authService, redisClient}
}

// handler category
// type categoryHandler struct {
// 	categoryService service.Service
// 	authService     config.Service
// }

// func NewHandlerCategory(categoryService service.Service, authService config.Service) *categoryHandler {
// 	return &categoryHandler{categoryService, authService}
// }
