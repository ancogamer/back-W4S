package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
	"w4s/security"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error3": err.Error()})
		return
	}
	if input.Password != "" {
		if err := security.VerifyPassword(user.Password, input.Password); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error1": err})
			return
		}
		if input.NewPassword != input.ConfirmNewPassword {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "A nova senha não conhecide !"})
			return
		}
		if err := models.PasswordCheck(input.NewPassword); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error2": err})
			return
		}
		password, err := models.BeforeSave(input.Password)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		input.Password = password
		if err := db.Model(&user).Update("password").Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"success": "senha trocada !"})
	}
	if input.Email != "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Não é possível trocar o email"})
		return
	}
	if err := db.Model(&user).Updates(input).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": user})
	return
}
