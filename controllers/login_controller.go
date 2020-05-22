package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"w4s/handlers"
	"w4s/models"
)

// Login is the signIn method
func Login(c *gin.Context) {
	var input models.LoginUser
	input.Token = c.Request.Header.Get("Authorization")
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Struct to store the data recovered from the database
	//Struct para armazenar os dados da base de dados
	login := models.User{
		Email:    input.Email,
		Nickname: input.Nickname,
		Password: input.Password,
	}
	err := login.Validate("login")
	//Validating the login inputs
	//Validando os inputs do login
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": err,
		})
		return
	}
	//Separating the where sql find, on other file, so i can use this func in other places if need.
	token := handlers.LoginFind(c, login, input)
	if token != "" {
		c.JSON(http.StatusOK, gin.H{"success": token})
	}
	return
}
