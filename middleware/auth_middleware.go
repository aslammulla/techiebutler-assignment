package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware is a middleware for authenticating requests
func AuthMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")

	if token != "12345" { // Replace this with actual token validation logic
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	c.Next()
}
