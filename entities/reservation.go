package entities

import (
	"time"
)

type Reservation struct {
	ID              uint      `gorm:"primaryKey"`
	Name            string    `gorm:"type:varchar(255);not null"`
	PhoneNumber     string    `gorm:"type:varchar(15);not null"`
	ReservationDate time.Time `gorm:"not null"`
	NumberOfPeople  int       `gorm:"not null"`
	Notes           string    `gorm:"type:text"`
	Status          string    `gorm:"type:enum('pending', 'confirmed', 'canceled');default:'pending'"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}
