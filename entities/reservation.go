package entities

import (
	"time"
)

type Reservation struct {
	ID     uint      `gorm:"primaryKey"`
	UserID uint      `gorm:"not null"`
	Date   time.Time `gorm:"not null"`
	Time   string    `gorm:"type:varchar(10);not null"`
	Guests int       `gorm:"not null"`
	Status string    `gorm:"type:enum('pending', 'confirmed', 'cancelled');default:'pending'"`
}
