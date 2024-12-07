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

	// Validasi input untuk memastikan tanggal, waktu, dan jumlah tamu valid
	if reservation.Date.IsZero() || reservation.Time == "" || reservation.Guests <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data, please provide valid date, time, and guest count."})
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

// Handler untuk mendapatkan semua reservasi (admin-only)
func (c *ReservationController) GetAllReservations(ctx *gin.Context) {
	reservations, err := c.ReservationModel.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}

	ctx.JSON(http.StatusOK, reservations)
}

// Handler untuk membatalkan reservasi
func (c *ReservationController) CancelReservation(ctx *gin.Context) {
	reservationID := ctx.Param("id")

	// Konversi reservationID ke uint
	uintReservationID, err := strconv.ParseUint(reservationID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}

	// Cek apakah pengguna yang melakukan request adalah admin atau pemilik reservasi
	// Untuk tujuan ini, kita asumsikan token sudah memuat informasi role dan userID
	userID := ctx.MustGet("userID").(uint)
	isAdmin := ctx.MustGet("role").(string) == "admin"

	// Jika bukan admin, pastikan user yang melakukan request adalah pemilik reservasi
	if !isAdmin && userID != uint(uintReservationID) {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to cancel this reservation"})
		return
	}

	// Batalkan reservasi
	if err := c.ReservationModel.Cancel(uint(uintReservationID)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel reservation"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}
