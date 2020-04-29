package main

import (
	"github.com/gin-gonic/gin"
	"w4s/DB"
	"w4s/controllers"
	"w4s/middleware"
)

//
/*type Authc struct {
	Email string `json:"email" binding:required `
	Token string `json:"token" binding:required `
}*/

func main() {
	//creating connection with database
	r := gin.Default()     //starting the gin. //Iniciando o gin
	db := DB.SetupModels() //Connection database //Conexão banco de dados
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	authorized := r.Group("/v1")
	r.POST("/login", controllers.Login)
	r.POST("/user/create", controllers.CreateUser)
	r.GET("/user/confirm", controllers.ConfirmUser)

	authorized.Use(middleware.AuthRequired)
	{
		authorized.GET("/searchall", controllers.FindUser)
		authorized.GET("/search", controllers.FindUserByNick)
		authorized.PATCH("/update/user", controllers.UpdateUser)
		authorized.PATCH("/logoff", controllers.Logoff)
		authorized.DELETE("/delete/user", controllers.SoftDeletedUserByNick)
		//sendo feitos
		/*authorized.PATCH("/update/user/email",controllers.UpdateUser)
		authorized.PATCH("/update/user/password",controllers.UpdateUser)*/
	}

	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") // listando e escutando no localhost:8080
	if err != nil {
		panic("NOT POSSIBLE RUN")
	}
}
