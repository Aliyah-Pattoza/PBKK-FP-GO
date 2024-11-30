package models

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type ReservationModel struct {
	DB *gorm.DB
}

func NewReservationModel(db *gorm.DB) *ReservationModel {
	return &ReservationModel{DB: db}
}

// fungsi untuk membuat reservasi baru
func (m *ReservationModel) Create(reservation *entities.Reservation) error {
	return m.DB.Create(reservation).Error
}

// fungsi untuk mendapatkan reservasi berdasarkan user ID
func (m *ReservationModel) GetByUserID(userID uint) ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	if err := m.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
