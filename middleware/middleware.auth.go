package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Authentication :
func Authentication(c *gin.Context) {
	authKey := c.GetHeader("Authorization")
	if authKey == "token2019wrong_token" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "token is invalid or already expire"})
		return
	}
	c.Next()
}
