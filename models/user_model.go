package models

import "time"

type User struct {
	ID        int `gorm:"primarykey"`
	Name      string
	Email     string
	Password  string
	Token     string
	Address   string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
