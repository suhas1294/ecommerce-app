package models

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Products []Product `json:"products"` // One-to-many: Category has many Products
}
