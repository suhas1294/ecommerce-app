package models

import "time"

type IssueReport struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UserID      uint      `json:"user_id"`
	Description string    `gorm:"type:text" json:"description"`
	Status      string    `gorm:"size:50;default:'Pending'" json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}
