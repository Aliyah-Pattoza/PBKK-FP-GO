package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserModel *models.UserModel
}

func NewUserController(userModel *models.UserModel) *UserController {
	return &UserController{UserModel: userModel}
}

// Handler untuk registrasi user baru
func (c *UserController) Register(ctx *gin.Context) {
	var user entities.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.UserModel.Create(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": user})
}

// Handler untuk mendapatkan user berdasarkan email
func (c *UserController) GetUserByEmail(ctx *gin.Context) {
	email := ctx.Param("email")
	user, err := c.UserModel.FindByEmail(email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
