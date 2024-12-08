package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type MenuController struct {
	MenuModel *models.MenuModel
}

func NewMenuController(menuModel *models.MenuModel) *MenuController {
	return &MenuController{MenuModel: menuModel}
}

// Middleware untuk memverifikasi role admin
func (c *MenuController) AdminOnly(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		ctx.Abort()
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !claims.VerifyIssuer("your-issuer", true) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		ctx.Abort()
		return
	}

	role, ok := claims["role"].(string)
	if !ok || role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin role required"})
		ctx.Abort()
		return
	}

	ctx.Next()
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
	c.AdminOnly(ctx)
	if ctx.IsAborted() {
		return
	}

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
	c.AdminOnly(ctx)
	if ctx.IsAborted() {
		return
	}

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
	c.AdminOnly(ctx)
	if ctx.IsAborted() {
		return
	}

	id := ctx.Param("id")
	if err := c.MenuModel.DeleteMenu(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
