package main

import (
	"../controllers/handler"
	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	r := gin.Default()

	r.GET("/login", handler.LoginCheck())
	/*router.POST("/somePost", posting)
	/router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	*/
	r.Run()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
