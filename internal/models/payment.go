package models

import "time"

type Payment struct {
	ID      uint      `gorm:"primaryKey" json:"id"`
	OrderID uint      `gorm:"uniqueIndex" json:"order_id"` // One-to-one mapping
	Method  string    `gorm:"size:50" json:"method"`
	Status  string    `gorm:"size:50" json:"status"`
	Amount  float64   `json:"amount"`
	PaidAt  time.Time `json:"paid_at"`
}
