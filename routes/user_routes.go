package routes

import (
	"pbkk-fp-go/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes sets up the routes for users
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
}
