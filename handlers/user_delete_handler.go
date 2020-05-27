package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	"w4s/models"
)

//Maria db treats false and true as tinyint, 0 for non deleted, 1 for deleted
func SoftDeletedUserByNick(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var user models.User
	var profile models.Profile
	if err := db.Where("nickname = ? AND deleted = ?", c.Query("nickname"), "0").Preload("Profile").First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
		return
	}
	//fmt.Println(user.Profile.IDUser)
	/*if err := db.Debug().Where("id_user = ?", user.Profile.IDUser).First(&profile).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}*/
	db.Model(&profile).Update("deleted_at", time.Now())
	db.Model(&user).Update(map[string]interface{}{"deleted": true, "deleted_at": time.Now(), "actived": false})
	c.JSON(http.StatusOK, gin.H{"success": "true"})
	return
}
