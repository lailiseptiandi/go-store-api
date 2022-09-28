package model

import "time"

type Product struct {
	ID           int       `json:"id"`
	CategoryID   int       `json:"category_id"`
	Name         string    `json:"name"`
	Slug         string    `json:"slug"`
	Descriptions string    `json:"descriptions"`
	Price        float64   `json:"price"`
	Quantity     int       `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
