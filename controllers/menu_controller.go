package controllers

import (
	"net/http"
	"pbkk-fp-go/entities"
	"pbkk-fp-go/models"
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"strings"
)

// MenuController mengatur logika terkait menu
type MenuController struct {
	MenuModel *models.MenuModel
}

// NewMenuController membuat instance baru dari MenuController
func NewMenuController(menuModel *models.MenuModel) *MenuController {
	return &MenuController{MenuModel: menuModel}
}

// Middleware untuk memverifikasi role admin
func (c *MenuController) AdminOnly(ctx *gin.Context) {
	// Ambil token dari header Authorization
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}

	// Token diawali dengan "Bearer "
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		ctx.Abort()
		return
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verifikasi signing method (gunakan secret key yang sesuai)
		return []byte("your-secret-key"), nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	// Dapatkan klaim (claims) dari token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !claims.VerifyIssuer("your-issuer", true) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		ctx.Abort()
		return
	}

	// Periksa apakah role user adalah 'admin'
	role := claims["role"].(string)
	if role != "admin" {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: Admin role required"})
		ctx.Abort()
		return
	}

	// Lanjutkan ke handler berikutnya
	ctx.Next()
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

// Handler untuk menambahkan menu baru (hanya admin)
func (c *MenuController) CreateMenu(ctx *gin.Context) {
	// Hanya admin yang bisa melakukan ini
	c.AdminOnly(ctx)
	if ctx.IsAborted() {
		return
	}

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

// Handler untuk memperbarui menu (hanya admin)
func (c *MenuController) UpdateMenu(ctx *gin.Context) {
	// Hanya admin yang bisa melakukan ini
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

	// Update menu
	if err := c.MenuModel.UpdateMenu(id, &menu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu updated successfully", "menu": menu})
}

// Handler untuk menghapus menu (hanya admin)
func (c *MenuController) DeleteMenu(ctx *gin.Context) {
	// Hanya admin yang bisa melakukan ini
	c.AdminOnly(ctx)
	if ctx.IsAborted() {
		return
	}

	id := ctx.Param("id")
	if err := c.MenuModel.DeleteMenu(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Menu deleted successfully"})
}
