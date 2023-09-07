package models

import "time"

type Permission struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RolePermission struct {
	ID           uint       `json:"id" gorm:"primarykey"`
	Role         string     `json:"role"`
	PermissionID uint       `json:"permission_id"`
	Permission   Permission `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
