package middlewares

import (
	"log"
	"net/http"
	"pbkk-fp-go/service"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware validates the JWT token from the Authorization header
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is empty")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is required"})
			c.Abort()
			return
		}

		claims, err := service.ValidateJWT(authHeader)
		if err != nil {
			log.Printf("Invalid token: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
