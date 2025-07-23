package models

import "time"

type Order struct {
	ID             uint                 `gorm:"primaryKey" json:"id"`
	UserID         uint                 `json:"user_id"` // Belongs to User
	User           User                 `json:"user"`
	TotalAmount    float64              `json:"total_amount"`
	PaymentStatus  string               `gorm:"size:50" json:"payment_status"`
	Payment        Payment              `json:"payment"`        // One-to-one: Order has one Payment
	OrderItems     []OrderItem          `json:"order_item"`     // One-to-many: Order has many OrderItems
	TrackingEvents []OrderTrackingEvent `json:"tracking_event"` // One-to-many: order tracking
	CreatedAt      time.Time            `json:"created_at"`
}
