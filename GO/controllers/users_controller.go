//POST /user
//Create a new user
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/authc"
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
		Deleted:  false,
		Token:    "",
	}
	err := user.Validate("createuser") //Validating the inputs/ Validando os inputs
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
	//Saving the new User on the database/ Salvando o novo usuario na base de dados
	if dbc := db.Create(&user); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
		return
	} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
	if err := SendConfirmationCreateAccountEmail(user.Email, c); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"waiting": "Verifique seu email !",
	})
	return
}
func ConfirmUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var UserToken models.AccountCreatedToken
	if err := db.Where("token = ?", c.Query("t")).First(&UserToken).Error; err != nil {
		UserToken.Token = c.Query("t")
		if err := db.Create(&UserToken).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if err := authc.ValidateToken(c.Query("t")); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var user models.User
		if err := db.Where("email = ?", c.Query("e")).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Registro não encontrado",
			})
			return
		}
		if err := db.Model(&user).Update("actived", true).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		//Here i'm returnin the message of account confirmated !/ Aqui estou retornando a mensagem que a conta foi confirmada ! :D
		//Future plains style with HTML and CSS, maybe some javascript too/ Planos futuros, estilizar com HTML e CSS, talvez um javascript junto
		c.JSON(http.StatusOK, gin.H{"success": "Conta confirmada ! "})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "link já utilizado"})
	return
}

func UpdateUser(c *gin.Context) {
	handlers.UpdateUser(c)
}
func FindUser(c *gin.Context) {
	handlers.FindUser(c)
}
func FindUserByNick(c *gin.Context) {
	handlers.FindUserByNick(c)
}
func SoftDeletedUserByNick(c *gin.Context) {
	handlers.SoftDeletedUserByNick(c)
}
