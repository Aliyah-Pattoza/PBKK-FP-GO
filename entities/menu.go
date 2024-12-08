package entities

import (
	"time"
)

type Menu struct {
	ID           uint      `gorm:"primaryKey"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text;not null"`
	Price        float64   `gorm:"type:decimal(8,2);not null"`
	Image        string    `gorm:"type:varchar(255)"`
	Category     string    `gorm:"type:enum('food', 'drink', 'dessert');not null"`
	Availability bool      `gorm:"default:true"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
