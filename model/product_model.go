package model

import "time"

type Product struct {
	ID         int
	CategoryID int
	Name       string
	Price      int
	Qty        int
	Images     string
	CreatedAt  time.Time
	UpdateAt   time.Time
}
