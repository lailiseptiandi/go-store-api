package models

import "time"

type User struct {
	ID        int    `gorm:"primarykey"`
	Email     string `gorm:"unique"`
	Name      string
	Password  string
	Token     string
	Address   string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
