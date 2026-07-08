package models

import "time"

type MenuItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`           // Buat storytelling "Menu Jujur"
	Price       int       `gorm:"not null" json:"price"` // Pakai integer (Rupiah) biar aman dari bug float
	IsAvailable bool      `gorm:"default:true" json:"is_available"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
