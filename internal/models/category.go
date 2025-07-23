package models

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `gorm:"uniqueIndex;size:100;not null" json:"name"`
	Products []Product `gorm:"foreignKey:CategoryRef" json:"products"` // One-to-many: Category has many Products
}

/*
Why Products field has a gorm tag foreignkey in it ?
it is to tell gorm - "Hey Gorm, CategoryRef is the FK in Product that points to me (Category)."
*/
