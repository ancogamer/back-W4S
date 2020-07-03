package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"w4s/controllers"
	"w4s/db"
	"w4s/middleware"
)

func main() {
	//creating connection with database
	r := gin.Default() //starting the gin. //Iniciando o gin

	DB := db.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", DB)
		c.Next()
	})

	//Loading HTML page and they needs(css,etc)
	r.Static("/css", "tela_alterar_senha/css")
	r.Static("/images", "tela_alterar_senha/images")
	r.LoadHTMLFiles("tela_alterar_senha/index.html")
	//Un Authorized Routes
	r.GET("", controllers.Ping)
	r.POST("/login", controllers.Login)

	r.POST("/create/user", controllers.CreateUser)
	r.POST("/create/user/resendlink", controllers.ResentCreateAccountLink)
	r.GET("/confirm/user", controllers.ConfirmUser)
	//User Recovery Password stuff
	r.POST("/user/password/recovery", controllers.RecoveryPasswordUser)
	recoveryPassword := r.Group("/user/password/recovery")
	//Uses a 2 middleware called AuthRequired2
	recoveryPassword.Use(middleware.AuthRequiredRecoveryPassword)
	{
		recoveryPassword.GET("", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})
		recoveryPassword.PUT("", controllers.ChangeExternalPassword)
	}

	//Normal Middleware
	authorized := r.Group("/v1")
	authorized2 := r.Group("/v2")
	authorized.Use(middleware.AuthRequired)
	{
		//Create the profile, only free, without check if the user already have a base profile
		//Cria o perfil, unica liberada, sem checar se o usuario já tem perfil base
		authorized.PATCH("/create/user/createprofile", controllers.CreateProfile) //Cria um perfil base
	}
	authorized2.Use(middleware.AuthRequired2)
	{
		//USER Links URL
		//encontra o perfil do usuario, busca pelo email e puxa junto o perfil do usuario
		//Find the user profile, search by the email and preload the user profile
		authorized2.GET("/searchall/user", controllers.FindAllUsers)        //Search all the users
		authorized2.GET("/search/user/profile", controllers.FindUserByNick) //Search the by nick, preload the user
		//Updates
		//Atualizações
		authorized2.PATCH("/update/user", controllers.UpdateUser) //Involves the User model, email or password
		//Logoff
		authorized2.PATCH("/logoff", controllers.Logoff) //Logoff from the system
		//Delete soft
		//Deletar soft
		authorized2.DELETE("/delete/user", controllers.SoftDeletedUserByNick) //Soft delete
		//Table Links URL
		authorized2.POST("/create/table", controllers.CreateTable) //Create table
		authorized2.GET("/searchall/user/table", controllers.FindAllUserTables)
		authorized2.GET("/searchall/table", controllers.FindAllTables) //Search by all the tables//Procura por todas as tabelas
		authorized2.GET("/searchone/table", controllers.FindOneTables) //Search one table//Procura por uma mesa

		authorized2.PUT("/update/table", controllers.UpdateTable)              //Update one table//Atualiza uma mesa
		authorized2.DELETE("/delete/table", controllers.DeleteTable)           //Delete one table//Deleta uma mesa
		authorized2.PATCH("/update/table/userjoin", controllers.UserJoinTable) //Join a user to the table//Coloca um usuario a mesa
	}
	//Run the ser
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}
