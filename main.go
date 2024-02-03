package main

import (
	"gin-mnc/controllers"
	"gin-mnc/initializers"

	"github.com/gin-gonic/gin"
)

func init(){
	
initializers.LoadEnvVariables()
initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.GET("/posts", controllers.PostsIndex)
	r.Run()
}