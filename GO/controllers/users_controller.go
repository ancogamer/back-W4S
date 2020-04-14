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

//Create User
func CreateUser(c *gin.Context) {
	fmt.Println("create FUNC")
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
		Token:    "",
	}
	err := user.Validate("") //Validating the inputs/ Validando os inputs
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Saving the new User on the database/ Salvando o novo usuario na base de dados
	if dbc := db.Create(&user); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
		return
	} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
	c.JSON(http.StatusOK, gin.H{"success": user})
	return
}
func UpdateUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("nickname = ?", c.Query("nickname")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	// Validate input
	var input models.UserInputUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Updates(input)
	c.JSON(http.StatusOK, gin.H{"success": user})
}

//Find all users on the database
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)
	if user != nil { //checking if something was returned.
		c.JSON(http.StatusOK, gin.H{
			"success": user,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"error": "nenhum registro encontrado",
	})
	return
}

//Find a user by his(her) nickname/Encontrando um usuario pelo seu nick(url)
func FindUserByNick(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.User
	if err := db.Where("nickname = ?", c.Query("nickname")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": user,
	})
}
