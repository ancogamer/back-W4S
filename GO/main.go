package main

import (
	"w4s/DB"
	"w4s/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	//creating connection with database
	r := gin.Default() //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	//Login
	r.POST("/login", controllers.Login)

	//Cria usuario
	r.POST("/user", controllers.CreateUser)

	//apiRoutes := main.Group("/api", auth.AuthorizeJWT()){

	//Start the routes group, that need authentication
	r.GET("/user", controllers.FindUser)


	//r.POST("/login", controllers.Login) //still in progress //ainda a ser feito

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
}
