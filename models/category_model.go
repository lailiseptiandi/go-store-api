package models

import (
	"gorm.io/gorm"
)

type Category struct {
	ID   int    `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	gorm.Model
}
