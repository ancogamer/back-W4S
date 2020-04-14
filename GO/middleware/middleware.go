package middleware

import (
	"github.com/gin-gonic/gin"
	"w4s/authc"
)

func AuthRequired(c *gin.Context) {
	authc.ValidateToken(c)
	c.Next()
	return
}
