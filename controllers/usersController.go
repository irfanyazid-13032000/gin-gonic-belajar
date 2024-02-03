package controllers

import (
	"gin-mnc/initializers"
	model "gin-mnc/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var bodyRequest struct {
		Email    string `json:"email"`
    Password string `json:"password"`
	}

	c.Bind(&bodyRequest)

	hash,err := bcrypt.GenerateFromPassword([]byte(bodyRequest.Password),10)

	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})
		return
	}

	


	user := model.User{
    Email:    bodyRequest.Email,
    Password: string(hash),
  }

	result := initializers.DB.Statement.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
    "user": user,
  })


}

