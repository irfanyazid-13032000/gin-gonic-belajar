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

func PostsShow(c *gin.Context) {
	var post model.Post
  id := c.Param("id")

  initializers.DB.Statement.DB.First(&post, id)

  c.JSON(200, gin.H{
    "post": post,
  })
}


func PostsUpdate(c *gin.Context) {
	var bodyRequest struct {
    Title string `json:"title"`
    Body  string `json:"body"`
  }

  c.Bind(&bodyRequest)

  id := c.Param("id")

  var post model.Post
  initializers.DB.Statement.DB.First(&post, id)

  post.Title = bodyRequest.Title
  post.Body = bodyRequest.Body

  result := initializers.DB.Statement.DB.Save(&post)

  if result.Error!= nil {
    c.Status(400)
    return
  }

  c.JSON(200, gin.H{
    "post": post,
  })
}


func PostsDelete(c *gin.Context) {
	var post model.Post
  id := c.Param("id")

  initializers.DB.Statement.DB.First(&post, id)

  result := initializers.DB.Statement.DB.Delete(&post)

  if result.Error!= nil {
    c.Status(400)
    return
  }

  c.JSON(200, gin.H{
    "post": post,
  })
}