package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
	"w4s/models"
	"w4s/security"
)

func LoginFind(c *gin.Context, login models.User, input models.LoginUser) string {
	db := c.MustGet("db").(*gorm.DB)
	//Checking by nickname
	if login.Email == "" {
		if err := db.Where("nickname = ?", input.Nickname).Find(&login).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error:": "Nickname ou senha incorretos",
			})
			return ""
		}
	} else {
		//Checking by email
		if err := db.Where("email = ?", input.Email).Find(&login).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error:": "Email ou senha incorretos",
			})
			return ""
		}
	}
	//(hashadpassword,password),
	//hashad = crypted password, password is the normal one/ hashadpassword = é a senha cryptografada, passoword é a senha normal
	if err := security.VerifyPassword(login.Password, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
		})
		return ""
	}
	token, err := authc.GenerateJWT(login, 24)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Não foi possível o acesso, tente mais tarde",
		})
		return ""
	}
	//Saving the new token on the user(Database)/ Salvando o novo token no usuario(Database)
	db.Model(login).Update("token", token)
	return token
}
