package routes

import (
	"assignment-3/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/*.html")
	router.Static("/assets", "./assets")
	router.GET("/", controllers.GetResult)

	return router
}
