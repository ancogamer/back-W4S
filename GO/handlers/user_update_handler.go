package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if input.Password != "" {
		changePassword(input.Password, c)
	}
	if input.Email != "" {
		changeEmail(input.Email, c)
	}
	if err := db.Model(&user).Updates(input).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{"success": user})
	return
}
func changePassword(password string, c *gin.Context) {

}
func changeEmail(email string, c *gin.Context) {

}
