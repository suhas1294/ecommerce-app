package models

import "time"

type Product struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	Price        float64   `gorm:"not null" json:"price"`
	Stock        int       `gorm:"default:0" json:"stock"`
	CategoryID   uint      `json:"category_id"` // Foreign key
	Category     Category  `json:"category"`    // Belongs to Category, we are having this field apart fomr category_id just to eager load data
	WishlistedBy []User    `gorm:"many2many:user_wishlist" json:"wishlisted_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
