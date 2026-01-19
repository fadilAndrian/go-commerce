package middleware

import (
	"net/http"
	"strings"

	"github.com/fadilAndrian/go-commerce/internal/helper"
	"github.com/gin-gonic/gin"
)

func NewAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")

		if header == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "Unauthorized",
			})
			c.Abort()
			return
		}

		headerPart := strings.Split(header, " ")
		if len(headerPart) != 2 || headerPart[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "Invalid token format",
			})
			c.Abort()
			return
		}

		token := headerPart[1]

		userId, err := helper.VerifyJWT(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"err": "Token is invalid or expired",
			})
			c.Abort()
			return
		}

		c.Set("userId", userId)

		c.Next()
	}
}
