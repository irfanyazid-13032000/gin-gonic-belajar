package controllers

import (
	"gin-mnc/initializers"
	model "gin-mnc/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
			"error": "failed to create user",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
    "user": user,
  })


}

func Login(c *gin.Context) {
	var bodyRequest struct {
			Email    string `json:"email"`
			Password string `json:"password"`
	}

	c.Bind(&bodyRequest)

	var user model.User

	result := initializers.DB.Where("email =?", bodyRequest.Email).First(&user)

	if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"error": "failed to login",
			})
			return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(bodyRequest.Password)); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"error": "failed to login",
			})
			return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":     user.ID,
			"expires": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
					"error": "failed to create token",
			})
			return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization",tokenString,3600 * 24 * 30,"","",false,true)

	c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
	})
}


