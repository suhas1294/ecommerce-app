package models

type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"user_id"`    // Belongs to Order
	ProductID uint    `json:"product_id"` // Belongs to Product
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
