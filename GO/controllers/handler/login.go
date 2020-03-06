package handler

import (
	"net/http"
	
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusOK, map[string]string{
			"Nicolas": "Bemvindo",
		})

	}
}
