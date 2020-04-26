//POST /user
//Create a new user
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	"w4s/handlers"
	"w4s/models"
)

//Create User
func CreateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	//Validating input
	var input models.UserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Creating user
	user := models.User{
		Nickname: input.Nickname,
		Email:    input.Email,
		Password: input.Password,
		Name:     input.Name,
		Lastname: input.Lastname,
		Actived:  false,
		Token:    "",
	}
	err := user.Validate("") //Validating the inputs/ Validando os inputs
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Password, err = models.BeforeSave(user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": err.Error(),
		})
		return
	}
	user.Created = time.Now().Unix()
	//Saving the new User on the database/ Salvando o novo usuario na base de dados
	if dbc := db.Create(&user); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
		return
	} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
	ConfirmationEmail(user.Email, c)
	c.JSON(http.StatusOK, gin.H{
		"waiting": "Verifique seu email !",
	})
	return
}
func ConfirmUserTOTP(c *gin.Context) {
	handlers.ConfirmUserTOTP1(c)
}
func UpdateUser(c *gin.Context) {
	handlers.UpdateUser1(c)
}
func FindUser(c *gin.Context) {
	handlers.FindUser1(c)
}
func FindUserByNick(c *gin.Context) {
	handlers.FindUserByNick1(c)
}
func SoftDeletedUserByNick(c *gin.Context) {
	handlers.SoftDeletedUserByNick1(c)
}
