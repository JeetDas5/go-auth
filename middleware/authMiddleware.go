package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	helper "github.com/jeetdas5/golang-jwt/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(token)

		if err != "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("first_name", claims.First_Name)
		c.Set("last_name", claims.Last_Name)
		c.Set("uid", claims.Uid)
		c.Set("user_type", claims.User_type)

		c.Next()
	}
}
