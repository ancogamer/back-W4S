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
		defer db.Close()
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
	err := user.Validate("")
	if err!=nil{
		defer db.Close()
		c.JSON(http.StatusNotImplemented, gin.H{
			"err":err,
		})
	}else{
		err :=user.BeforeSave()
		if err!=nil{
			c.JSON(http.StatusConflict, gin.H{"erro":err})
		}else {
			if dbc := db.Create(&user); dbc.Error != nil { //return the error by JSON
				defer db.Close()
				c.JSON(http.StatusBadRequest, gin.H{"erro": dbc.Error})
			} else { //return the post data if is ok, by JSON
				defer db.Close()
				c.JSON(http.StatusOK, gin.H{"data": user})
			}
		}
	}
}
//Find all users on the database
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	defer db.Close()
	if user!=nil{//checking if something was returned.
		c.JSON(http.StatusNotFound, gin.H{
			"error":"nenhum registro encontrado",
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"user":user,
		})
	}
}