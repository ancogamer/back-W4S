package handlers

import (
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
	var user models.User
	var profile models.Profile
	//db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)
	if err := db.Table("users").Select("*").Joins("inner join profiles on profiles.id = profile_id").
		Where("nickname = ?", c.Query("nickname")).Scan(&profile).Scan(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	/*
		if err := db.Debug().Preload("Profile").First(&userProfile, "id_user = ?", c.Query("nickname")).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Registro não encontrado",
			})
			return
		}
	*/
	user.Profile = profile

	c.JSON(http.StatusOK, gin.H{
		"success": user,
	})
}
