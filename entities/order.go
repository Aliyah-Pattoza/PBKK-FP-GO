package entities

import "time"

type Order struct {
	ID         uint        `gorm:"primaryKey"`
	UserID     uint        `gorm:"not null"`
	TotalPrice float64     `gorm:"type:decimal(8,2);not null"`
	Status     string      `gorm:"type:enum('pending', 'confirmed', 'canceled');default:'pending'"`
	OrderDate  time.Time   `gorm:"not null;autoCreateTime"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	User       User        `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time   `gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `gorm:"autoUpdateTime"`
}
