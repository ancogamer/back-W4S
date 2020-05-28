package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
	"w4s/models"
)

//Normal Middleware
func AuthRequired(c *gin.Context) {
	authc.ValidateLoginToken(c)
	c.Next()
	return
}

//Recovery PasswordMiddleware
func AuthRequired2(c *gin.Context) {
	tokenCheck(c)
	c.Next()
	return
}

//Check if the user created a base profile
func AuthRequired3(c *gin.Context) {
	claim := authc.ValidateLoginToken(c)
	if claim != c.Query("e") {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "internal server error"})
		return
	}
	c.Next()
	return
}

//Cheking Token
func tokenCheck(c *gin.Context) {
	var token models.UserAccountBadListToken
	db := c.MustGet("db").(*gorm.DB)
	token.Token = c.Query("t")
	if db.Where("token = ?", token.Token).First(&token).RecordNotFound() {
		if _, err := authc.ValidateToken(token.Token); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Alguma coisa não deu certo, por favor, requiste novamente a recuperação de senha"})
			return
		}
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Link já utilizado !"})
	return
}
