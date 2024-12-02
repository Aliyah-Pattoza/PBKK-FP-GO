package middlewares

import (
	"log"
	"net/http"
	"pbkk-fp-go/helpers"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware memeriksa apakah pengguna memiliki peran yang sesuai
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil klaim dari context
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Konversi klaim ke tipe *helpers.Claims
		userClaims, ok := claims.(*helpers.Claims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Log untuk debugging
		log.Printf("User Role: %s, Required Role: %s", userClaims.Role, requiredRole)

		// Periksa apakah role pengguna sesuai
		if userClaims.Role != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		// Lanjutkan ke handler berikutnya jika role sesuai
		c.Next()
	}
}
