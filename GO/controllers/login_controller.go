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
	Email    string `json:"email" binding:required `
	Password string `json:"password" binding:required`
}

// Login is the signIn method
func Login(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		defer db.Close()
		return
	}
	login := models.User{}

	if err := db.Where("email = ?", input.Email).Find(&login); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro:": "Email ou senha incorretos",
		})
		defer db.Close()
		return
	}
	if err := security.VerifyPassword(login.Password, input.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
		})
		defer db.Close()
		return
	}
	token,err:= authc.GenerateJWT(login)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"erro":"Não foi possível criar um token de acesso, tente mais tarde",
		})
		return
	}
	c.JSON(http.StatusOK,token)
	return
}
