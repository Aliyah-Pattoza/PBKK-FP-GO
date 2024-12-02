package controllers

import (
	"log"
	"net/http"
	"pbkk-fp-go/config"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/helpers"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input entities.User

	// Bind JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println("Error Binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("Parsed Input:", input)

	// Hash password sebelum disimpan
	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		log.Println("Error Hashing Password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	input.Password = hashedPassword

	// Simpan user ke database
	if err := config.DB.Create(&input).Error; err != nil {
		log.Println("Error Saving User to Database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save user"})
		return
	}

	log.Println("User Registered Successfully:", input)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully!"})
}

func Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var user entities.User

	// Bind input JSON ke struct
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cari user berdasarkan email
	if err := config.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Validasi password
	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate JWT dengan role pengguna
	token, err := helpers.GenerateJWT(user.Email, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
