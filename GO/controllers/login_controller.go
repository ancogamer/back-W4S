package controllers

import (
	"net/http"
	"w4s/authc"
	"w4s/models"
	"w4s/security"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//
type LoginUser struct {
	Email    string `json:"email" binding:"required" `
	Password string `json:"password" binding:"required"`
}

// Login is the signIn method
func Login(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Struct to store the data recovered from the database /Struct para armazenar os dados da base de dados
	login := models.User{
		Email:    input.Email,
		Password: input.Password,
	}
	err := login.Validate("login") //Validating the login inputs / Validando os inputs do login
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": err,
		})
		return
	}
	if err := db.Where("email = ?", input.Email).Find(&login).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error:": "Email ou senha incorretos",
		})
		return
	}
	//(hashadpassword,password),
	//hashad = crypted password, password is the normal one/ hashadpassword = é a senha cryptografada, passoword é a senha normal
	if err := security.VerifyPassword(login.Password, input.Password); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
		})
		return
	}
	token, err := authc.GenerateJWT(login)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Não foi possível o acesso, tente mais tarde",
		})
		return
	}
	//Saving the new token on the user(Database)/ Salvando o novo token no usuario(Database)
	db.Model(login).Update("token", token)
	c.JSON(http.StatusOK, gin.H{"success": token})
	return
}
