package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
)

func Logoff(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	/*//Validating the Token / Validando o token
	authc.ValidateToken(c)
	//If is ok do the next/ Se estiver tudo certo, continua*/
	usertoken := models.LogoffListTokens{
		Token: c.Request.Header.Get("Authorization"),
	}
	if dbc := db.Create(&usertoken); dbc.Error != nil { //Return the error by JSON / Retornando o erro por JSON
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": dbc.Error})
		return
	} //Return the post data if is ok, by JSON/ Retornando o que foi postado se tudo ocorreu certo
	c.JSON(http.StatusOK, gin.H{"success": ""})
}
