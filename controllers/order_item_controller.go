package controllers

import (
	"net/http"
	"strconv"

	"pbkk-fp-go/entities"
	"pbkk-fp-go/repository"

	"github.com/gin-gonic/gin"
)

type OrderItemController struct {
	Repository repository.OrderItemRepository
}

func (ctrl *OrderItemController) GetAll(c *gin.Context) {
	orderItems, err := ctrl.Repository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItems)
}

func (ctrl *OrderItemController) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	orderItem, err := ctrl.Repository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "OrderItem not found"})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func (ctrl *OrderItemController) Create(c *gin.Context) {
	var orderItem entities.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.Repository.Create(&orderItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, orderItem)
}

func (ctrl *OrderItemController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var orderItem entities.OrderItem
	if err := c.ShouldBindJSON(&orderItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	orderItem.ID = uint(id)
	if err := ctrl.Repository.Update(&orderItem); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderItem)
}

func (ctrl *OrderItemController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := ctrl.Repository.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
