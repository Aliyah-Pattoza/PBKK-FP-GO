package entities

type OrderItem struct {
	ID       uint    `gorm:"primaryKey"`
	OrderID  uint    `gorm:"not null"`
	MenuID   uint    `gorm:"not null"`
	Quantity int     `gorm:"not null"`
	Subtotal float64 `gorm:"not null"`
}
