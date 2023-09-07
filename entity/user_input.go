package entity

import "github.com/lailiseptiandi/golang-toko-online/models"

type (
	RegisterUserInput struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Roles    string `json:"roles"`
		Address  string `json:"address"`
	}

	LoginInput struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	CheckEmailUser struct {
		Email string `json:"email" binding:"required"`
	}

	GetUserID struct {
		ID   int `uri:"id" binding:"required"`
		User models.User
	}
)
