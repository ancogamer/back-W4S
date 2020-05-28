package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
	"w4s/models"
	"w4s/security"
)

func LoginFind(c *gin.Context, user models.User, input models.LoginUser) string {

	db := c.MustGet("db").(*gorm.DB)
	//Checking by email
	if err := db.Where("email = ? ", input.Email).Find(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error:": "Não\nencontramos sua ficha em local algum, por favor, nos dê\ncredenciais válidas, ou vá fazer seu registro com o\nRegistrador.”",
		})
		return ""
	}
	if user.Actived == false {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Ative sua conta pelo link enviado no email !",
		})
		return ""
	}
	//(hashadpassword,password),
	//hashad = crypted password, password is the normal one/ hashadpassword = é a senha cryptografada, passoword é a senha normal
	if err := security.VerifyPassword(user.Password, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Não\nencontramos sua ficha em local algum, por favor, nos dê\ncredenciais válidas, ou vá fazer seu registro com o\nRegistrador.”",
		})
		return ""
	}
	token, err := authc.GenerateJWT(user.Email, 86400)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Não foi possível o acesso, tente mais tarde",
		})
		return ""
	}
	//Saving the new token on the user(Database)/ Salvando o novo token no usuario(Database)
	db.Model(user).Update("token", token)

	return token
}
