package middlewares

import (
	"log"
	"net/http"
	"pbkk-fp-go/service"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware ensures the user has the correct role
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			log.Println("Claims not found in context")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		userClaims, ok := claims.(*service.Claims)
		if !ok || userClaims.Role != requiredRole {
			log.Printf("Access denied: Required Role: %s, User Role: %s", requiredRole, userClaims.Role)
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
