Still think about this
/*package service

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"w4s/models"
)

type LoginService interface {
	Login(username string, password string) bool
}
func (controller *loginController) Login(c *gin.Context) string{
	db:= c.MustGet("db").(*gorm.DB)
	user := models.User{}
	err:=c.ShouldBindJSON(&user)
	if err!=nil{
		defer db.Close()
		return ""
	}
	isAuthenticated:=controller.loginService.Login(user.Email,user.Password)
	if isAuthenticated{
		defer db.Close()
		return controller.jwtService.GerenateToken(user.Email,true)
	}
	defer db.Close()
	return ""
}*/