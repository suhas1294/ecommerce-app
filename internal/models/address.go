package models

type Address struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	UserID  uint   `json:"user_id"` // Belongs to User
	Line1   string `gorm:"size:255" json:"line_1"`
	City    string `gorm:"size:100" json:"city"`
	Pincode string `gorm:"size:20" json:"pincode"`
	Country string `gorm:"size:50" json:"country"`
}
