package controllers

import (
	"gin-mnc/initializers"
	model "gin-mnc/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	var bodyRequest struct {
		Title string `json:"title"`
    Body  string `json:"body"`
	}

	c.Bind(&bodyRequest)


	post := model.Post{Title: bodyRequest.Title, Body: bodyRequest.Body}
	result := initializers.DB.Statement.DB.Create(&post) 

	if result.Error != nil {
		c.Status(400)
		return
	}


	c.JSON(200, gin.H{
		"post": post,
	})
}


func PostsIndex(c *gin.Context) {
	var posts []model.Post
  initializers.DB.Statement.DB.Find(&posts)

  c.JSON(200, gin.H{
    "posts": posts,
  })
}