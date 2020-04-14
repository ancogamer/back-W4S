package controllers

import (
	"net/http"
	"w4s/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateProfileInput struct {
	Avatar         string `json:"avatar" binding:"required"`         //longtext no BD (mysql-MariaDB)
	DataNascimento string `json:"datanascimento" binding:"required"` //maximo 8 digitos
	Nickname       string `json:"nickname" binding:"required"`
}

func CreateProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input CreateProfileInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//Creating Profile
	profile := models.Profile{
		Avatar:         input.Avatar,
		DataNascimento: input.DataNascimento,
		Nickname:       input.Nickname,
	}
	if dbc := db.Create(&profile); dbc.Error != nil { //return the error by JSON
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": dbc.Error})
		return
	}
	//return the post data if is ok, by JSON
	c.JSON(http.StatusOK, gin.H{"success": profile})
	return

}

//Find all users on the database
func FindProfile(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var profile []models.Profile
	db.Find(&profile)
	if profile != nil { //checking if something was returned.
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "nenhum registro encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": profile,
	})
	return

}
