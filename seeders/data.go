package seeders

import (
	"log"
	"time"

	"pbkk-fp-go/config"
	"pbkk-fp-go/entities"
)

func SeedDatabase() {
	db := config.DB

	// Seeder for Users
	users := []entities.User{
		{Name: "John Doe", Email: "john@example.com", Password: "password123", Role: "user"},
		{Name: "Jane Admin", Email: "admin@example.com", Password: "adminpassword", Role: "admin"},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Println("Failed to seed users:", err)
	}

	// Seeder for Menus
	menus := []entities.Menu{
		{Name: "Espresso", Description: "A shot of rich coffee", Price: 3.50, Category: "drink", Availability: true},
		{Name: "Tiramisu", Description: "Classic Italian dessert", Price: 5.00, Category: "dessert", Availability: true},
		{Name: "Chicken Sandwich", Description: "Grilled chicken with fresh vegetables", Price: 6.00, Category: "food", Availability: true},
	}
	if err := db.Create(&menus).Error; err != nil {
		log.Println("Failed to seed menus:", err)
	}

	// Seeder for Reservations
	reservations := []entities.Reservation{
		{Name: "Alice", PhoneNumber: "1234567890", ReservationDate: time.Now().Add(24 * time.Hour), NumberOfPeople: 2, Notes: "Near window", Status: "confirmed"},
		{Name: "Bob", PhoneNumber: "0987654321", ReservationDate: time.Now().Add(48 * time.Hour), NumberOfPeople: 4, Notes: "Birthday party", Status: "pending"},
	}
	if err := db.Create(&reservations).Error; err != nil {
		log.Println("Failed to seed reservations:", err)
	}

	// Seeder for Orders and Order Items
	orders := []entities.Order{
		{UserID: 1, TotalPrice: 9.50, Status: "confirmed"},
		{UserID: 2, TotalPrice: 15.00, Status: "pending"},
	}
	if err := db.Create(&orders).Error; err != nil {
		log.Println("Failed to seed orders:", err)
	}

	orderItems := []entities.OrderItem{
		{OrderID: 1, MenuID: 1, Quantity: 1, Price: 3.50},
		{OrderID: 1, MenuID: 2, Quantity: 1, Price: 5.00},
		{OrderID: 2, MenuID: 3, Quantity: 2, Price: 6.00},
	}
	if err := db.Create(&orderItems).Error; err != nil {
		log.Println("Failed to seed order items:", err)
	}

	log.Println("Database seeded successfully!")
}
