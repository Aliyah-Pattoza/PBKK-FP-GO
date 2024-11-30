package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	OrderModel *models.OrderModel
}

func NewOrderController(orderModel *models.OrderModel) *OrderController {
	return &OrderController{OrderModel: orderModel}
}

// Handler untuk membuat pesanan baru
func (c *OrderController) CreateOrder(ctx *gin.Context) {
	var order entities.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.OrderModel.Create(&order); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Order created successfully", "order": order})
}

// Handler untuk mendapatkan pesanan berdasarkan ID
func (c *OrderController) GetOrderByID(ctx *gin.Context) {
	id := ctx.Param("id")

	// Konversi string ID ke uint
	uintID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}

	order, err := c.OrderModel.GetByID(uint(uintID)) // Konversi berhasil
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
