package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"w4s/DB"
	"w4s/authc"
	"w4s/controllers"
	"w4s/models"
)
//
type Authc struct {
	Email string `json:"email" binding:required `
	Token string `json:"token" binding:required `
}
func AuthRequired(c *gin.Context)  {
	var input Authc
		fmt.Println(input.Token)
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error":err.Error(),
		})
		return
	}
		// We want to make sure the token is set, bail if not
		if input.Token ==""  {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Usuario":"não logado",
			})
			return
		}

		//Fill up the Struct User model /Preenchendo o model User
		user := models.User{Email:input.Email,
			Token:input.Token,
		}
		err:=user.Validate("login")
		if err!=nil{
			c.JSON(http.StatusConflict,err)
		}
		authc.ValidateToken(c)

		return
		c.Next()
}

func main() {
	//creating connection with database
	r := gin.Default() //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	authorized := r.Group("/v2")

	// AuthRequired() middleware just in the "authorized" group.
	r.POST("/login", controllers.Login)
	//Cria usuario
	r.POST("/user", controllers.CreateUser)


	authorized.Use(AuthRequired)
	{
		authorized.GET("/user", controllers.FindUser)
		authorized.GET("/seach/:nickname", controllers.FindUserByNick)
	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
}
