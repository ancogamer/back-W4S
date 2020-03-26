package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"w4s/DB"
	"w4s/controllers"
)
func AuthRequired(c *gin.Context)  {
		userToken:= c.Param("token")

		// We want to make sure the token is set, bail if not
		if userToken ==""  {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Usuario":"não logado",
			})
			return
		}
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
	//Login
	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	authorized := r.Group("/v2")
	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	r.POST("/login", controllers.Login)
	//Cria usuario
	r.POST("/user", controllers.CreateUser)
	authorized.Use(AuthRequired)
	{
		authorized.GET("/user", controllers.FindUser)
	}


	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
}
