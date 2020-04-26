package handlers

import (
	"github.com/gin-gonic/gin"
	"w4s/authc"
)

func ConfirmUserTOTP1(c *gin.Context) {
	authc.TOTPValidation(c)
}
