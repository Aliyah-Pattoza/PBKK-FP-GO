package entities

import (
	"time"
)

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`
	TotalPrice float64     `gorm:"not null"`
	Status     string      `gorm:"type:enum('pending', 'completed', 'cancelled');default:'pending'"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	Items      []OrderItem `gorm:"foreignKey:OrderID"`
}
