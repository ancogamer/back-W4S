package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

//Maria DB treats false and true as tinyint, 0 for non deleted, 1 for deleted
func SoftDeletedUserByNick1(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	if err := db.Where("nickname = ? AND deleted = ?", c.Query("nickname"), "0").First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	db.Model(&user).Update("deleted", true)
	c.JSON(http.StatusOK, gin.H{"success": "true"})
}
