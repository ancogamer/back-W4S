package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

func FindUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var user []models.User
	db.Find(&user)

	c.JSON(http.StatusOK, gin.H{
		"user":user,
	})

}
