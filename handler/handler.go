package handler

import (
	"github.com/lailiseptiandi/go-store-api/config"
	"github.com/lailiseptiandi/go-store-api/service"
)

type globalHandler struct {
	globalService service.Service
	authService   config.Service
}

func NewHandler(globalService service.Service, authService config.Service) *globalHandler {
	return &globalHandler{globalService, authService}
}

// handler category
// type categoryHandler struct {
// 	categoryService service.Service
// 	authService     config.Service
// }

// func NewHandlerCategory(categoryService service.Service, authService config.Service) *categoryHandler {
// 	return &categoryHandler{categoryService, authService}
// }
