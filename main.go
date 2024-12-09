package main

import (
	"log"
	"pbkk-fp-go/config"
	"pbkk-fp-go/controllers"
	"pbkk-fp-go/middlewares"
	"pbkk-fp-go/models"
	"pbkk-fp-go/routes"
	"pbkk-fp-go/seeders"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDatabase()
	seeders.SeedDatabase()

	DB := config.DB

	// Inisialisasi router
	r := gin.Default()

	routes.RegisterOrderItemRoutes(r, DB)

	// Inisialisasi model dan controller
	menuModel := models.NewMenuModel(DB)
	menuController := controllers.NewMenuController(menuModel)

	orderModel := models.NewOrderModel(DB)
	orderController := controllers.NewOrderController(orderModel)

	reservationModel := models.NewReservationModel(DB)
	reservationController := controllers.NewReservationController(reservationModel)

	// Rute tanpa proteksi (public routes)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)
	r.GET("/menus", menuController.GetMenus) // Rute untuk mendapatkan daftar menu (public)

	// Rute dengan proteksi JWT
	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware()) // Middleware untuk validasi JWT
	{
		// Rute menu (admin only)
		menuRoutes := protected.Group("/menus")
		{
			menuRoutes.POST("/", middlewares.RoleMiddleware("admin"), menuController.CreateMenu)
			menuRoutes.PUT("/:id", middlewares.RoleMiddleware("admin"), menuController.UpdateMenu)
			menuRoutes.DELETE("/:id", middlewares.RoleMiddleware("admin"), menuController.DeleteMenu)
		}

		// Rute reservasi
		reservationRoutes := protected.Group("/reservations")
		{
			reservationRoutes.POST("/", reservationController.CreateReservation)
			reservationRoutes.GET("/", reservationController.GetReservations)
			reservationRoutes.DELETE("/:id", reservationController.CancelReservation)
			reservationRoutes.PUT("/:id", middlewares.RoleMiddleware("admin"), reservationController.UpdateReservationStatus)
		}

		// Rute pesanan
		orderRoutes := protected.Group("/orders")
		{
			orderRoutes.POST("/", orderController.CreateOrder)
			orderRoutes.GET("/:id", orderController.GetOrderByID)
		}

		// Rute profil pengguna
		protected.GET("/profile", controllers.Profile)

		// Rute khusus admin
		protected.GET("/admin", middlewares.RoleMiddleware("admin"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, admin!"})
		})

		// Rute khusus user
		protected.GET("/user", middlewares.RoleMiddleware("user"), func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome, user!"})
		})
	}

	// Jalankan server pada port 8080
	log.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
