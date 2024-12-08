package routes

import (
	"pbkk-fp-go/controllers"
	"pbkk-fp-go/middlewares"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes sets up the routes for users
// func RegisterUserRoutes(router *gin.Engine) {
// 	router.POST("/register", controllers.Register)
// 	router.POST("/login", controllers.Login)

// }

func RegisterUserRoutes(router *gin.Engine) {
	// Group protected routes
	protected := router.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware()) // Middleware untuk validasi JWT
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
}
