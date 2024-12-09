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
		{Name: "John Doe", Email: "john@gmail.com", Password: "password123", Role: "user"},
		{Name: "Jane Admin", Email: "admin@gmail.com", Password: "adminpassword", Role: "admin"},
	}
	if err := db.Create(&users).Error; err != nil {
		log.Println("Failed to seed users:", err)
		return
	}
	log.Println("Users seeded successfully")

	// Seeder for Menus
	menus := []entities.Menu{
		{Name: "Espresso", Description: "A shot of rich coffee", Price: 3.50, Category: "drink", Availability: true},
		{Name: "Tiramisu", Description: "Classic Italian dessert", Price: 5.00, Category: "dessert", Availability: true},
		{Name: "Chicken Sandwich", Description: "Grilled chicken with fresh vegetables", Price: 6.00, Category: "food", Availability: true},
	}
	if err := db.Create(&menus).Error; err != nil {
		log.Println("Failed to seed menus:", err)
		return
	}
	log.Println("Menus seeded successfully")

	// Seeder for Orders
	orders := []entities.Order{
		{UserID: users[0].ID, TotalPrice: 9.50, Status: "confirmed"},
		{UserID: users[1].ID, TotalPrice: 15.00, Status: "pending"},
	}
	if err := db.Create(&orders).Error; err != nil {
		log.Println("Failed to seed orders:", err)
		return
	}
	log.Println("Orders seeded successfully")

	// Seeder for Order Items
	orderItems := []entities.OrderItem{
		{OrderID: orders[0].ID, MenuID: menus[0].ID, Quantity: 1, Price: 3.50},
		{OrderID: orders[0].ID, MenuID: menus[1].ID, Quantity: 1, Price: 5.00},
		{OrderID: orders[1].ID, MenuID: menus[2].ID, Quantity: 2, Price: 6.00},
	}
	if err := db.Create(&orderItems).Error; err != nil {
		log.Println("Failed to seed order items:", err)
		return
	}
	log.Println("Order Items seeded successfully")

	// Seeder for Reservations
	reservations := []entities.Reservation{
		{
			UserID:          users[0].ID,
			Name:            "Alice",
			PhoneNumber:     "1234567890",
			ReservationDate: time.Now().Add(24 * time.Hour),
			NumberOfPeople:  2,
			Notes:           "Near window",
			Status:          "confirmed",
		},
		{
			UserID:          users[0].ID,
			Name:            "Bob",
			PhoneNumber:     "0987654321",
			ReservationDate: time.Now().Add(48 * time.Hour),
			NumberOfPeople:  4,
			Notes:           "Birthday party",
			Status:          "pending",
		},
	}
	if err := db.Create(&reservations).Error; err != nil {
		log.Println("Failed to seed reservations:", err)
		return
	}
	log.Println("Reservations seeded successfully")

	log.Println("Database seeded successfully!")
}
