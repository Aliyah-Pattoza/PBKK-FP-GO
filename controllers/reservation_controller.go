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

	// Set userID dari JWT
	userID := ctx.MustGet("userID").(uint)
	reservation.UserID = userID

	// Validasi tanggal dan jumlah orang
	if reservation.ReservationDate.Before(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Reservation date cannot be in the past"})
		return
	}
	if reservation.NumberOfPeople <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Number of guests must be greater than zero"})
		return
	}

	// Simpan reservasi
	if err := c.ReservationModel.Create(&reservation); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "Reservation created successfully",
		"reservation": reservation,
	})
}

// GetReservations handles fetching reservations based on user role
func (c *ReservationController) GetReservations(ctx *gin.Context) {
	userID := ctx.MustGet("userID").(uint)
	role := ctx.MustGet("role").(string)

	var reservations []entities.Reservation
	var err error

	if role == "admin" {
		// Admin dapat melihat semua reservasi
		reservations, err = c.ReservationModel.GetAll()
	} else {
		// User hanya dapat melihat reservasi miliknya sendiri
		reservations, err = c.ReservationModel.GetByUserID(userID)
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// CancelReservation handles cancelling a reservation
func (c *ReservationController) CancelReservation(ctx *gin.Context) {
	// Ambil ID reservasi
	reservationID := ctx.Param("id")
	uintReservationID, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	// Ambil userID dan role dari JWT
	userID := ctx.MustGet("userID").(uint)
	role := ctx.MustGet("role").(string)

	// Ambil reservasi berdasarkan ID
	reservation, err := c.ReservationModel.GetByID(uint(uintReservationID))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	// Periksa apakah user adalah admin atau pemilik reservasi
	if role != "admin" && reservation.UserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to cancel this reservation"})
		return
	}

	// Batalkan reservasi
	if err := c.ReservationModel.Cancel(uint(uintReservationID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel reservation: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}

// UpdateReservationStatus handles updating the status of a reservation (admin-only)
func (c *ReservationController) UpdateReservationStatus(ctx *gin.Context) {
	// Ambil ID reservasi dari parameter URL
	reservationID := ctx.Param("id")
	uintReservationID, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	// Ambil status baru dari body permintaan
	var requestBody struct {
		Status string `json:"status"`
	}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	// Validasi status baru
	validStatuses := map[string]bool{
		"pending":   true,
		"completed": true,
		"cancelled": true,
	}
	if !validStatuses[requestBody.Status] {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid status. Allowed statuses: pending, completed, cancelled"})
		return
	}

	// Perbarui status reservasi
	if err := c.ReservationModel.UpdateStatus(uint(uintReservationID), requestBody.Status); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reservation status: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation status updated successfully"})
}
