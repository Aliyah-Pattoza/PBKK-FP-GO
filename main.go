package main

import (
	"net/http"
	"pbkk-fp-go/config"
	"pbkk-fp-go/controllers/homepage"
)

func main() {
	// Inisialisasi koneksi database
	config.ConnectDatabase()

	// Inisialisasi router
	//r := gin.Default()

	// Homepage
	http.HandleFunc("/", homepage.Welcome)

	//Menu

	// Rute tanpa proteksi (public routes)
	//r.POST("/register", controllers.Register)
	//r.POST("/login", controllers.Login)

	// Rute dengan proteksi JWT
	//	protected := r.Group("/api")
	/*	protected.Use(middlewares.JWTAuthMiddleware()) // Middleware untuk validasi JWT
		{
			// Rute untuk semua pengguna dengan token valid
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
		//r.Run(":8080")*/
	http.ListenAndServe(":8080", nil)
}
