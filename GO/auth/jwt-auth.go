package auth

import (
	"github.com/gin-gonic/gin"
	"w4s/service"
)

func AuthorizeJWT() gin.HandlerFunc{
	return func(c *gin.Context){
		const BEARE_SCHEMA= "Bearer"
		authHeader:=c.GetHeader("Authorization")
		tokenString :=authHeader[len(BEARE_SCHEMA):]
		token,err:=service.JWTService().ValidateToken(tokenString)
	}
}

