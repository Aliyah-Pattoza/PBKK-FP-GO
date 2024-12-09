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

func (c *MenuController) GetMenus(ctx *gin.Context) {
	menus, err := c.MenuModel.GetAllMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menus: " + err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"menus": menus})
}

func (c *MenuController) CreateMenu(ctx *gin.Context) {
	var menu entities.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi tambahan
	if menu.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than or equal to 0"})
		return
	}
	if menu.Category != "food" && menu.Category != "drink" && menu.Category != "dessert" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	if err := c.MenuModel.Create(&menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Menu created successfully", "menu": menu})
}

func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	var menu entities.Menu
	if err := ctx.ShouldBindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi tambahan
	if menu.Price < 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Price must be greater than or equal to 0"})
		return
	}
	if menu.Category != "food" && menu.Category != "drink" && menu.Category != "dessert" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category"})
		return
	}

	if err := c.MenuModel.UpdateMenu(id, &menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully", "menu": menu})
}

func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.MenuModel.DeleteMenu(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
