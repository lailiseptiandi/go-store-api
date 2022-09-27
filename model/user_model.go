package model

import "time"

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Token     string
	Address   string
	Roles     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
