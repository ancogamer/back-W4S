package authc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pquerna/otp/totp"
	"net/http"
	"time"
	"w4s/models"
	"w4s/utility"
)

func TOTPGenerate(email string, c *gin.Context) string {
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "Find A Table System",
		AccountName: email,
		Period:      14400,
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err})
	}
	// Now Validate that the user's successfully added the passcode.
	passcode := utility.Randomsequence()
	fmt.Println(passcode)
	valid := totp.Validate(passcode, key.Secret())
	fmt.Println(valid)
	if valid {
		db := c.MustGet("db").(*gorm.DB)
		UserTotp := models.TOTPkey{
			Key:       key.Secret(),
			UserEmail: email,
			Actived:   time.Now().Unix(),
		}
		dbc := db.Create(&UserTotp)
		if dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
		}
		c.JSON(http.StatusOK, gin.H{"success": dbc.Error})
		return passcode
	}
	return ""
}
func TOTPValidation(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var secretKey models.TOTPkey
	passcode := c.Query("passcode")
	if err := db.Where("email = ? ", c.Query("email")).First(&secretKey).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": "Registro não encontrado",
		})
	}
	valid := totp.Validate(passcode, secretKey.Key)
	if valid {
		var user models.User
		if err := db.Where("email = ?", c.Query("email")).First(&user).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error": "Registro não encontrado",
			})
		}
		db.Model(&user).Update("actived", true)
		c.JSON(http.StatusOK, gin.H{"success": "true"})
	}

}
