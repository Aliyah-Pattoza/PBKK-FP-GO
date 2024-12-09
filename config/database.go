package config

import (
	"log"
	"pbkk-fp-go/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/pbkk-fp-go?parseTime=true"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// AutoMigrate all entities
	err = db.AutoMigrate(
		&entities.User{},
		&entities.Menu{},
		&entities.Order{},
		&entities.OrderItem{},
		&entities.Reservation{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = db
}
