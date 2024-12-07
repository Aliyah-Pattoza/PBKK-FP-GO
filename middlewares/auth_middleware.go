package middlewares

import (
	"net/http"
	"pbkk-fp-go/helpers"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware validates the JWT token from the Authorization header
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil token dari header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		// Validasi token
		claims, err := helpers.ValidateJWT(authHeader)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Simpan informasi klaim ke context untuk digunakan di handler lain
		c.Set("claims", claims)
		c.Next()
	}
}
