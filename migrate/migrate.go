package main

import (
	"gin-mnc/initializers"
	"gin-mnc/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&model.Post{})
	initializers.DB.AutoMigrate(&model.User{})
	initializers.DB.AutoMigrate(&model.Customer{})
	initializers.DB.AutoMigrate(&model.Transaction{})
}