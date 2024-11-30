package entities

type Menu struct {
	ID       uint    `gorm:"primaryKey"`
	Name     string  `gorm:"type:varchar(100);not null"`
	Category string  `gorm:"type:enum('food', 'drink', 'dessert');not null"`
	Price    float64 `gorm:"not null"`
	Stock    int     `gorm:"default:0"`
}
