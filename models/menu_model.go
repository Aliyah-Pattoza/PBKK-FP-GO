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

func (m *MenuModel) GetAllMenus() ([]entities.Menu, error) {
	var menus []entities.Menu
	if err := m.DB.Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}

func (m *MenuModel) Create(menu *entities.Menu) error {
	return m.DB.Create(menu).Error
}

func (m *MenuModel) UpdateMenu(id string, menu *entities.Menu) error {
	return m.DB.Model(&entities.Menu{}).Where("id = ?", id).Updates(menu).Error
}

func (m *MenuModel) DeleteMenu(id string) error {
	return m.DB.Where("id = ?", id).Delete(&entities.Menu{}).Error
}
