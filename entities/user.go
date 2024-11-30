package entities

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"type:varchar(100);unique;not null"`
	Password string `gorm:"not null"`
	Role     string `gorm:"type:enum('user', 'admin');default:'user'"`
}
