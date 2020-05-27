package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

//Find all users on the database
func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var users []models.User

	if err := db.Where("deleted = ? AND actived = ?", "0", true).Preload("Profile").Preload("Tables").Find(&users).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Nenhum registro encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": users,
	})
	return
}

//Find a user by his(her) nickname/Encontrando um usuario pelo seu nick(url)
func FindUserByNick(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user models.Profile
	if err := db.Debug().Where("nickname = ? AND deleted = ?", c.Query("e"), "0").Preload("Profile.User").First(&user).Error; err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": user,
	})
}
