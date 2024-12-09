package models

import (
	"errors"
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
	if err := m.DB.Create(reservation).Error; err != nil {
		return errors.New("failed to create reservation: " + err.Error())
	}
	return nil
}

// Fungsi untuk mendapatkan reservasi berdasarkan user ID
func (m *ReservationModel) GetByUserID(userID uint) ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	if err := m.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, errors.New("failed to retrieve reservations for user: " + err.Error())
	}
	return reservations, nil
}

// Fungsi untuk mendapatkan semua reservasi (admin-only)
func (m *ReservationModel) GetAll() ([]entities.Reservation, error) {
	var reservations []entities.Reservation
	if err := m.DB.Find(&reservations).Error; err != nil {
		return nil, errors.New("failed to retrieve all reservations: " + err.Error())
	}
	return reservations, nil
}

// Fungsi untuk mendapatkan reservasi berdasarkan ID
func (m *ReservationModel) GetByID(reservationID uint) (*entities.Reservation, error) {
	var reservation entities.Reservation
	if err := m.DB.First(&reservation, reservationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("reservation not found")
		}
		return nil, errors.New("failed to retrieve reservation: " + err.Error())
	}
	return &reservation, nil
}

// Fungsi untuk membatalkan reservasi
func (m *ReservationModel) Cancel(reservationID uint) error {
	reservation, err := m.GetByID(reservationID)
	if err != nil {
		return err
	}
	// Perbarui status reservasi menjadi "cancelled"
	reservation.Status = "cancelled"
	if err := m.DB.Save(reservation).Error; err != nil {
		return errors.New("failed to cancel reservation: " + err.Error())
	}
	return nil
}

// UpdateStatus updates the status of a reservation
func (m *ReservationModel) UpdateStatus(reservationID uint, status string) error {
	// Cari reservasi berdasarkan ID
	var reservation entities.Reservation
	if err := m.DB.First(&reservation, reservationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("reservation not found")
		}
		return errors.New("failed to find reservation: " + err.Error())
	}

	// Perbarui status reservasi
	reservation.Status = status
	if err := m.DB.Save(&reservation).Error; err != nil {
		return errors.New("failed to update reservation status: " + err.Error())
	}
	return nil
}
