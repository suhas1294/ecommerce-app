package models

import "time"

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Email        string    `gorm:"uniqueIndex;size:100;not null" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"password_hash"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	Orders       []Order   `json:"orders"`                                  // One-to-many: User has many Orders
	Wishlist     []Product `gorm:"many2many:user_wishlist" json:"wishlist"` // Many-to-many: Wishlist
	Addresses    []Address `json:"addresses"`                               // One-to-many: User has multiple addresses
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// i dont have to add any gorm tag for Orders field since gorm will understand automatically - convention over configuration,
// in case we want to be explicit, we can give : Orders []Order `gorm:"foreignKey:UserID"`
