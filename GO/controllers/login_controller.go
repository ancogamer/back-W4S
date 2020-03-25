package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"w4s/models"
	"w4s/service"
)

type loginController struct {
	Email    string `json:"email" binding:required `
	Password string `json:"password" binding:required`
}
// Login is the signIn method
func (controller *loginController) Login(c *gin.Context) {
	var input models.Login
	if err:= c.ShouldBindJSON(&input);err !=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
	login:=models.Login{
		Email:    input.Email,
		Password: input.Password,
	}
	if err:=db.Where("email= ?", c.Params("email")).First(&login).Error; err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Email não encontrado"})
		return
	}
	//isAuthenticated:=service.LoginService.Login()
	//if isAuthenticated{
	//	defer db.Close()
	//	return controller.jwtService.GerenateToken(user.Email,true)
	//}
	//defer db.Close()
	//return ""
}

