package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
	"w4s/models"
)

func CreateTable(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var input models.TableInput
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user models.User
	if err := db.Where("nickname = ?", c.Query("nickname")).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "não encontrado o nickname",
		})
		return
	}
	table := models.Table{
		Name:                 input.Name,
		Description:          input.Description,
		NumberOfParticipants: 1,
		Thumbnail:            input.Thumbnail,
		AdventureLink:        input.AdventureLink,
	}
	//REVER POR CAUSA DO USUARIO
	if err := db.Create(&table).Error; err != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	db.Model(table).Association("Users").Append([]*models.User{&user})
	insertPictures(c, table.ID)
	c.JSON(http.StatusOK, gin.H{"success": "table created"})
}
func FindAllTables(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tables []models.Table

	if err := db.Preload("User").Preload("User.Profile").Find(&tables).Error; err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Nenhum registro encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": tables,
	})
	return
}
func insertPictures(c *gin.Context, TableId uint) {
	db := c.MustGet("db").(*gorm.DB)
	var pictures models.Picture
	if err := c.BindJSON(pictures); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
	}

	pictures.TableID = TableId
	split := strings.Split(pictures.PictureFile, " ")
	for i := 0; i < len(split); i++ {
		if err := db.Create(&pictures).Error; err != nil { //Return the error by JSON / Retornando o erro por JSON
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	return
}
