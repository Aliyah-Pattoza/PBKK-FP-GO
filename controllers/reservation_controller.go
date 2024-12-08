package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ReservationController struct {
	ReservationModel *models.ReservationModel
}

func NewReservationController(reservationModel *models.ReservationModel) *ReservationController {
	return &ReservationController{ReservationModel: reservationModel}
}

// CreateReservation handles creating a new reservation
func (c *ReservationController) CreateReservation(ctx *gin.Context) {
	var reservation entities.Reservation
	if err := ctx.ShouldBindJSON(&reservation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validate date and time
	if reservation.ReservationDate.Before(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Reservation date cannot be in the past"})
		return
	}
	if reservation.NumberOfPeople <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Number of guests must be greater than zero"})
		return
	}

	// Save reservation
	if err := c.ReservationModel.Create(&reservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "Reservation created successfully",
		"reservation": reservation,
	})
}

// GetReservationsByUserID handles fetching reservations by user ID
func (c *ReservationController) GetReservationsByUserID(ctx *gin.Context) {
	userID := ctx.Param("userID")

	// Convert userID from string to uint
	uintUserID, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Fetch reservations
	reservations, err := c.ReservationModel.GetByUserID(uint(uintUserID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations: " + err.Error()})
		return
	}

	if len(reservations) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "No reservations found for this user"})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// GetAllReservations handles fetching all reservations (admin-only)
func (c *ReservationController) GetAllReservations(ctx *gin.Context) {
	// Check admin role
	isAdmin := ctx.MustGet("role").(string) == "admin"
	if !isAdmin {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to view all reservations"})
		return
	}

	// Fetch all reservations
	reservations, err := c.ReservationModel.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// CancelReservation handles canceling a reservation
func (c *ReservationController) CancelReservation(ctx *gin.Context) {
	// Retrieve reservation ID
	reservationID := ctx.Param("id")
	uintReservationID, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	// Retrieve userID and role from context
	userID := ctx.MustGet("userID").(uint)
	isAdmin := ctx.MustGet("role").(string) == "admin"

	// Retrieve the reservation by ID
	reservation, err := c.ReservationModel.GetByID(uint(uintReservationID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	// Check if user is admin or the owner of the reservation
	if !isAdmin && reservation.ID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to cancel this reservation"})
		return
	}

	// Cancel the reservation
	if err := c.ReservationModel.Cancel(uint(uintReservationID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel reservation"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}
