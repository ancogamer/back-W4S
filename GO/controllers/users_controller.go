//POST /user
//Create a new user
package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)
//Struct from the inputs
type UserInput struct{
	Nickname string  `json:"nickname" binding:required`
	Email 	 string  `json:"email" binding:required `
	Password string  `json:"password" binding:required`
	Name     string  `json:"name" binding:required`
	Lastname string  `json:"lastname" binding:required`
}


//Create User
func CreateUser( c *gin.Context){
	fmt.Println("create FUNC")
	db:= c.MustGet("db").(*gorm.DB)
	//Validating input
	var input UserInput

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
		Token:"",
	}
	err := user.Validate("") //Validating the inputs/ Validando os inputs
	if err!=nil{
		defer db.Close()
		c.JSON(http.StatusNotImplemented, gin.H{
			"err":err,
		})
		return
	}
	user.Token = ""
	//Saving the new User on the database/ Salvando o novo usuario na base de dados
	if dbc := db.Create(&user); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
		c.JSON(http.StatusBadRequest, gin.H{"erro": dbc.Error})
		return
	} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
		c.JSON(http.StatusOK, gin.H{"data": user})
	return

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
		return
	}
		c.JSON(http.StatusOK, gin.H{
			"user":user,
		})
	return
}
//Find a user by his(her) nickname/Encontrando um usuario pelo seu nick(url)
func FindUserByNick(c *gin.Context){
	db:=c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("nickname = ?", c.Param("nickname")).First(&user).Error;err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Record not found!",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"User":user,
	})

}
