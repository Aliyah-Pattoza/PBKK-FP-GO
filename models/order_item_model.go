package models

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type OrderItemModel struct {
	DB *gorm.DB
}

func NewOrderItemModel(db *gorm.DB) *OrderItemModel {
	return &OrderItemModel{DB: db}
}

// fungsi untuk membuat item pesanan
func (m *OrderItemModel) Create(item *entities.OrderItem) error {
	return m.DB.Create(item).Error
}
