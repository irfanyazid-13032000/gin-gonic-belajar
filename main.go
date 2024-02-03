package main

import (
	"gin-mnc/controllers"
	"gin-mnc/initializers"
	"gin-mnc/middleware"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

	// Set logrus to write to a file
	logFile, err := os.OpenFile("gin-mnc.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatal("Error opening log file: ", err)
	}
	logrus.SetOutput(logFile)
	logrus.SetFormatter(&logrus.JSONFormatter{})
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
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.POST("/customers", controllers.CustomerCreate)

	r.POST("/transactions", controllers.CreateTransaction)

	r.Run()
}
