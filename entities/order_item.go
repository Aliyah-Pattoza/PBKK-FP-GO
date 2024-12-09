package entities

import "time"

type OrderItem struct {
	ID        uint      `gorm:"primaryKey"`
	OrderID   uint      `gorm:"not null"`
	MenuID    uint      `gorm:"not null"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"type:decimal(8,2);not null"`
	Order     Order     `gorm:"foreignKey:OrderID"`
	Menu      Menu      `gorm:"foreignKey:MenuID"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
