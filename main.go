package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"w4s/db"
	"w4s/controllers"
	"w4s/middleware"
)

func main() {
	//creating connection with database
	r := gin.Default()     //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	//Loading HTML page and they needs(css,etc)
	r.Static("/css", "tela_alterar_senha/css")
	r.Static("/images", "tela_alterar_senha/images")
	r.LoadHTMLFiles("tela_alterar_senha/index.html")

	//Un Authorized Routes
	r.POST("/login", controllers.Login)
	r.POST("/create/user", controllers.CreateUser)
	r.GET("/confirm/user", controllers.ConfirmUser)
	//User Recovery Password stuff
	r.POST("/user/password/recovery", controllers.RecoveryPasswordUser)
	recoveryPassword := r.Group("/user/password/recovery")
	//Uses a 2 middleware called AuthRequired2
	recoveryPassword.Use(middleware.AuthRequired2)
	{
		recoveryPassword.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		recoveryPassword.PUT("", controllers.ChangeExternalPassword)
	}

	//Normal Middleware
	authorized := r.Group("/v1")
	authorized.Use(middleware.AuthRequired)
	{
		//USER Links URL
		authorized.GET("/searchall/user", controllers.FindAllUsers)
		authorized.GET("/search/user", controllers.FindUserByNick)
		authorized.PATCH("/create/user/createprofile", controllers.CreateProfile)
		authorized.PATCH("/update/user", controllers.UpdateUser)
		authorized.PATCH("/logoff", controllers.Logoff)
		authorized.DELETE("/delete/user", controllers.SoftDeletedUserByNick)
		//Table Links URL
		authorized.POST("/create/table", controllers.CreateTable)
		authorized.GET("/searchall/table", controllers.FindAllTables)

	}

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}
