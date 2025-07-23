package models

import "time"

type Product struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Name         string    `gorm:"size:100;not null" json:"name"`
	Description  string    `gorm:"type:text" json:"description"`
	Price        float64   `gorm:"not null" json:"price"`
	Stock        int       `gorm:"default:0" json:"stock"`
	CategoryRef  uint      `json:"category_ref"`                           // foreign key - breaking convention over configuration
	Category     Category  `gorm:"foreignKey:CategoryRef" json:"category"` // Belongs to Category, we are having this field apart fomr category_id just to eager load data
	WishlistedBy []User    `gorm:"many2many:user_wishlist" json:"wishlisted_by"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

/*
BREAKING CONVENTION OVER CONFIGURATION ;

1. The table has a column category_ref.
2. That column is a foreign key pointing to categories.id
3. GORM will use category_ref when preloading Product.Category
4. You can query or preload like: `db.Preload("Category").Find(&products)`

This is equivalent of  :

CREATE TABLE products (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    category_ref INTEGER,
    price REAL NOT NULL,
    FOREIGN KEY (category_ref) REFERENCES categories(id)
);
------------------------------------

Eager loading due to Category field in product struct  :
GORM won’t attempt to load a Category struct because there’s no Category Category field.

We can add 'not null' in 'CategoryRef' assuming Category field is not there.
We can even add `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` to CategoryRef field.

In that case struct definition would be :
type Product struct {
    ID          uint   `gorm:"primaryKey"`
    Name        string `gorm:"size:100;not null"`
    CategoryRef uint   `gorm:"column:category_ref;not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
You don’t need to specify foreignKey:Category here because there’s no actual struct (Category)field.

--------------------------
*/
