package models

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type MenuModel struct {
	DB *gorm.DB
}

func NewMenuModel(db *gorm.DB) *MenuModel {
	return &MenuModel{DB: db}
}

// fungsi untuk mendapatkan semua menu
func (m *MenuModel) GetAllMenus() ([]entities.Menu, error) {
	var menus []entities.Menu
	if err := m.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

// fungsi untuk menambahkan menu baru
func (m *MenuModel) Create(menu *entities.Menu) error {
	return m.DB.Create(menu).Error
}
