package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"

	"github.com/gin-gonic/gin"
)

type MenuController struct {
	MenuModel *models.MenuModel
}

func NewMenuController(menuModel *models.MenuModel) *MenuController {
	return &MenuController{MenuModel: menuModel}
}

// Handler untuk mendapatkan semua menu
func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := c.MenuModel.GetAllMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus"})
		return
	}

	ctx.JSON(http.StatusOK, menus)
}

// Handler untuk menambahkan menu baru
func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu entities.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.MenuModel.Create(&menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Menu created successfully", "menu": menu})
}
