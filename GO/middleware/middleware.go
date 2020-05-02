package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
	"w4s/models"
)

func AuthRequired(c *gin.Context) {
	authc.ValidateLoginToken(c)
	c.Next()
	return
}
func AuthRequired2(c *gin.Context) {
	tokenCheck(c)
	c.Next()
	return
}
func tokenCheck(c *gin.Context) {
	var token models.UserAccountBadListToken
	db := c.MustGet("db").(*gorm.DB)
	token.Token = c.Query("t")
	if db.Where("token = ?", token.Token).First(&token).RecordNotFound() {
		if err := authc.ValidateToken(token.Token); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Alguma coisa não deu certo, por favor, requiste novamente a recuperação de senha"})
			return
		}
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Link já utilizado !"})
	return
}
