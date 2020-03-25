package service

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"w4s/models"
)


type JWTService interface{
	GenerateToken(user models.User) string
	ValidateToken(c *gin.Context) (*jwt.Token,error)
}

