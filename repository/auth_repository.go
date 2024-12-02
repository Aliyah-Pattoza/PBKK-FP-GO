package repository

import (
	"pbkk-fp-go/entities"

	"gorm.io/gorm"
)

type AuthRepository interface {
	EmailExist(email string) bool
	Register(req *entities.User) error
	GetUserByEmail(email string) (*entities.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entities.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *authRepository) EmailExist(email string) bool {
	var user entities.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}

func (r *authRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User

	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}
