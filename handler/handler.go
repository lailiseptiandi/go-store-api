package handler

import (
	"github.com/lailiseptiandi/golang-toko-online/config"
	"github.com/lailiseptiandi/golang-toko-online/service"
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
