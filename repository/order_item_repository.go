package repository

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func (r *OrderItemRepository) FindAll() ([]entities.OrderItem, error) {
	var orderItems []entities.OrderItem
	if err := r.DB.Find(&orderItems).Error; err != nil {
		return nil, err
	}
	return orderItems, nil
}

func (r *OrderItemRepository) FindByID(id int) (*entities.OrderItem, error) {
	var orderItem entities.OrderItem
	if err := r.DB.First(&orderItem, id).Error; err != nil {
		return nil, err
	}
	return &orderItem, nil
}

func (r *OrderItemRepository) Create(orderItem *entities.OrderItem) error {
	return r.DB.Create(orderItem).Error
}

func (r *OrderItemRepository) Update(orderItem *entities.OrderItem) error {
	return r.DB.Save(orderItem).Error
}

func (r *OrderItemRepository) Delete(id int) error {
	return r.DB.Delete(&entities.OrderItem{}, id).Error
}
