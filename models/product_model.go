package models

import (
	"gorm.io/gorm"
)

type Product struct {
	ID           int     `json:"id" gorm:"primarykey"`
	CategoryID   int     `json:"category_id"`
	Name         string  `json:"name"`
	Slug         string  `json:"slug"`
	Descriptions string  `json:"descriptions"`
	Price        float64 `json:"price"`
	Quantity     int     `json:"quantity"`
	UserID       uint    `json:"user_id"`
	User         User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	gorm.Model
}

type ProductImage struct {
	ID        int    `gorm:"primarykey"`
	ImageName string `json:"image_name"`
	ProductID int    `json:"product_id"`
	gorm.Model
}
