package models

import "time"

type OrderTrackingEvent struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	OrderID   uint      `json:"order_id"` // Belongs to Order
	Status    string    `gorm:"size:100" json:"status"`
	Location  string    `gorm:"size:100" json:"location"`
	EventTime time.Time `json:"event_time"`
}
