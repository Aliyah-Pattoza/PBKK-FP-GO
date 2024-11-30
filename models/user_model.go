package models

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{DB: db}
}

// fungsi untuk mencari user berdasarkan email
func (m *UserModel) FindByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := m.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// fungsi untuk membuat user baru
func (m *UserModel) Create(user *entities.User) error {
	return m.DB.Create(user).Error
}
