package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	ReservationModel *models.ReservationModel
}

func NewReservationController(reservationModel *models.ReservationModel) *ReservationController {
	return &ReservationController{ReservationModel: reservationModel}
}

// Handler untuk membuat reservasi baru
func (c *ReservationController) CreateReservation(ctx *gin.Context) {
	var reservation entities.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ReservationModel.Create(&reservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully", "reservation": reservation})
}

// Handler untuk mendapatkan reservasi berdasarkan user ID
func (c *ReservationController) GetReservationsByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	// Konversi string userID ke uint
	uintUserID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	reservations, err := c.ReservationModel.GetByUserID(uint(uintUserID)) // Konversi berhasil
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No reservations found"})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}
