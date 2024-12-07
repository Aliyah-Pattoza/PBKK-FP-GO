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

// Fungsi untuk membuat reservasi baru
func (m *ReservationModel) Create(reservation *entities.Reservation) error {
	return m.DB.Create(reservation).Error
}

// Fungsi untuk mendapatkan reservasi berdasarkan user ID
func (m *ReservationModel) GetByUserID(userID uint) ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	if err := m.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

// Fungsi untuk mendapatkan semua reservasi (admin only)
func (m *ReservationModel) GetAll() ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	if err := m.DB.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

// Fungsi untuk membatalkan reservasi
func (m *ReservationModel) Cancel(reservationID uint) error {
	var reservation entities.Reservation
	if err := m.DB.First(&reservation, reservationID).Error; err != nil {
		return err
	}
	reservation.Status = "cancelled"
	return m.DB.Save(&reservation).Error
}
