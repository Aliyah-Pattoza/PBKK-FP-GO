package main

import (
	"pbkk-fp-go/config"
	"pbkk-fp-go/controllers"
	"pbkk-fp-go/middlewares"
	"pbkk-fp-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDatabase()

	// Inisialisasi router
	r := gin.Default()

	// Inisialisasi model untuk menu dan reservasi
	menuModel := models.NewMenuModel(config.DB)
	reservationModel := models.NewReservationModel(config.DB)

	// Inisialisasi controller untuk menu dan reservasi
	menuController := controllers.NewMenuController(menuModel)
	reservationController := controllers.NewReservationController(reservationModel)

	// Rute tanpa proteksi (public routes)
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Rute dengan proteksi JWT
	protected := r.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware()) // Middleware untuk validasi JWT
	{
		// Rute untuk profil pengguna dengan token valid
		protected.GET("/profile", controllers.Profile)

		// Rute untuk Welcome messages berdasarkan role
		protected.GET("/welcome", func(c *gin.Context) {
			// Ambil role dari JWT token yang sudah divalidasi
			userRole := c.MustGet("role").(string)

			if userRole == "admin" {
				c.JSON(200, gin.H{"message": "Welcome, Admin!"})
			} else if userRole == "user" {
				c.JSON(200, gin.H{"message": "Welcome, User!"})
			} else {
				c.JSON(403, gin.H{"error": "Unauthorized role"})
			}
		})

		// Rute untuk Menu - Semua pengguna bisa melihat menu
		protected.GET("/menu", menuController.GetMenus)

		// Rute untuk admin menambah menu
		protected.POST("/menu", middlewares.RoleMiddleware("admin"), menuController.CreateMenu)

		// Rute untuk admin mengedit menu
		protected.PUT("/menu/:id", middlewares.RoleMiddleware("admin"), menuController.UpdateMenu)

		// Rute untuk admin menghapus menu
		protected.DELETE("/menu/:id", middlewares.RoleMiddleware("admin"), menuController.DeleteMenu)

		// Rute untuk Reservasi
		protected.POST("/reservations", reservationController.CreateReservation)

		// Rute untuk mendapatkan reservasi berdasarkan userID
		protected.GET("/reservations/:userID", reservationController.GetReservationsByUserID)

		// Rute untuk membatalkan reservasi (user atau admin)
		protected.PUT("/reservations/:id/cancel", middlewares.RoleMiddleware("admin"), reservationController.CancelReservation)
		protected.PUT("/reservations/:id/cancel", reservationController.CancelReservation)
	}

	// Jalankan server pada port 8080
	r.Run(":8080")
}
