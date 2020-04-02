package controllers

import (
	"fmt"
	"net/http"
	"w4s/authc"
	"w4s/models"
	"w4s/security"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//
type LoginUser struct {
	Email 	 string  `json:"email" binding:required `
	Password string  `json:"password" binding:required`
}

// Login is the signIn method
func Login(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var input LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Struct to store the data recovered from the database /Struct para armazenar os dados da base de dados
	login := models.User{
		Email:     input.Email,
		Password:  input.Password,
	}
	err := login.Validate("login") //Validating the login inputs / Validando os inputs do login
	if err!=nil {
		c.JSON(http.StatusNotImplemented, gin.H{
			"err": err,
		})
		db.Close()
		return
	}
	fmt.Println(login.Password,input.Password)
	if err:= db.Where("email = ?", input.Email).Find(&login).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro:": "Email ou senha incorretos",
		})
		return
	}
	//(hashadpassword,password),
	//hashad = crypted password, password is the normal one/ hashadpassword = é a senha cryptografada, passoword é a senha normal
	if err := security.VerifyPassword(login.Password, input.Password); err != nil {
		fmt.Println(login.Password,input.Password)
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
		})
		return
	}
	token,err:= authc.GenerateJWT(login)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"erro":"Não foi possível o acesso, tente mais tarde",
		})
		return
	}
	//Saving the new token on the user(Database)/ Salvando o novo token no usuario(Database)
	db.Model(login).Update("token",token)
	c.JSON(http.StatusOK,token)

	return
}
