package controllers

import (
	"log"
	"net/http"
	"pbkk-fp-go/config"
	"pbkk-fp-go/dto"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/helper"
	"pbkk-fp-go/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input dto.RegisterRequest

	// Bind input JSON ke DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Printf("Error Binding JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi input
	if input.Password != input.PasswordConfirmation {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password confirmation does not match"})
		return
	}

	hashedPassword, err := helper.HashPassword(input.Password)
	if err != nil {
		log.Printf("Error Hashing Password: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Persiapan data user baru
	newUser := entities.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	// Simpan ke database
	if err := config.DB.Create(&newUser).Error; err != nil {
		log.Printf("Error Saving User to Database: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	log.Println("User Registered Successfully:", newUser)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

func Login(c *gin.Context) {
	var input dto.LoginRequest
	var user entities.User

	// Bind input JSON ke DTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validasi input
	if input.Email == "" || input.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and password are required"})
		return
	}

	// Cari user berdasarkan email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Verifikasi password
	if err := helper.VerifyPassword(user.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Buat token JWT
	token, err := service.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	response := dto.LoginResponse{
		ID:    int(user.ID),
		Name:  user.Name,
		Token: token,
	}
	c.JSON(http.StatusOK, response)
}
