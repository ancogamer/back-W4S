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
	/*Email    string `json:"email" binding:required `
	Password string `json:"password" binding:required`*/

	Email 	 string  `json:"email" binding:required `
	Password string  `json:"password" binding:required`

}

// Login is the signIn method
func Login(c *gin.Context){
	db := c.MustGet("db").(*gorm.DB)
	var input LoginUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorrt": err.Error(),
		})
		return
	}
	//struct to store the data recovered from the database
	login := models.User{
		Email:     input.Email,
		Password:  input.Password,
	}
	err := login.Validate("login")
	if err!=nil {
		defer db.Close()
		c.JSON(http.StatusNotImplemented, gin.H{
			"err": err,
		})
		return
	}
	fmt.Println(login.Password,input.Password)
	if err:= db.Where("email = ?", input.Email).Find(&login).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"erro:": "Email ou senha incorretos",
		})
		return
	}
	//(password,hashadpassword)
	if err := security.VerifyPassword(input.Password, login.Password); err != nil {
		fmt.Println(login.Password,input.Password)
		fmt.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "senha incorreta",
			"errocode":err,
		})
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
