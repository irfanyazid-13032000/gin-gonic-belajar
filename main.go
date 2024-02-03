package main

import (
	"gin-mnc/controllers"
	"gin-mnc/initializers"
	"gin-mnc/middleware"

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
	r.GET("/posts/:id", controllers.PostsShow)
	r.PUT("/posts/:id", controllers.PostsUpdate)
	r.DELETE("/posts/:id", controllers.PostsDelete)


	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate",middleware.RequireAuth, controllers.Validate)
	r.Run()
}