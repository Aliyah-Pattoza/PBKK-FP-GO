package models

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type OrderModel struct {
	DB *gorm.DB
}

func NewOrderModel(db *gorm.DB) *OrderModel {
	return &OrderModel{DB: db}
}

// fungsi untuk membuat pesanan baru
func (m *OrderModel) Create(order *entities.Order) error {
	return m.DB.Create(order).Error
}

// fungsi untuk mendapatkan pesanan berdasarkan ID
func (m *OrderModel) GetByID(id uint) (*entities.Order, error) {
	var order entities.Order
	if err := m.DB.Preload("Items").Where("id = ?", id).First(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}
