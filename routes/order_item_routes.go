package routes

import (
	"pbkk-fp-go/controllers"
	"pbkk-fp-go/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RegisterOrderItemRoutes sets up the routes for order items
func RegisterOrderItemRoutes(router *gin.Engine, db *gorm.DB) {
	// Initialize repository and controller
	orderItemRepo := repository.OrderItemRepository{DB: db}
	orderItemCtrl := controllers.OrderItemController{Repository: orderItemRepo}

	// Define the group of routes
	orderItems := router.Group("/order-items")
	{
		orderItems.GET("/", orderItemCtrl.GetAll)
		orderItems.GET("/:id", orderItemCtrl.GetByID)
		orderItems.POST("/", orderItemCtrl.Create)
		orderItems.PUT("/:id", orderItemCtrl.Update)
		orderItems.DELETE("/:id", orderItemCtrl.Delete)
	}
}
