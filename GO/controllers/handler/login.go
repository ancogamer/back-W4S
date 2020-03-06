package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]string{
			"Hello": "Word",
		})

	}
}
