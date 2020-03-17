//POST /user
//Create a new user
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

type CreateUserInput struct{
	Nickname string  `json:"nickname" binding:required`
	Email 	 string  `json:"email" binding:required `
	Password string  `json:"password" binding:required`
	Name     string  `json:"name" binding:required`
	Lastname string  `json:"lastname" binding:required`
}
func CreateUser( c *gin.Context){
	db:= c.MustGet("db").(*gorm.DB)
	//Validating input
	var input CreateUserInput

	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	//Creating user
	user:=models.User{
		Nickname: input.Nickname,
		Email:input.Email,
		Password :input.Password,
		Name: input.Name,
		Lastname: input.Lastname,

	}
	if dbc:= db.Create(&user); dbc.Error != nil {//return the error by JSON
		c.JSON(http.StatusBadRequest,gin.H{"erro":dbc.Error})
	}else { //return the post data if is ok, by JSON
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}
//Find all users on the database
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"user":user,
	})
}