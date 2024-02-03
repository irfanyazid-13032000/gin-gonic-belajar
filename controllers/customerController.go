package controllers

import (
	"gin-mnc/initializers"
	model "gin-mnc/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomerCreate(c *gin.Context) {
	// Binding request body to bodyRequest struct
	var bodyRequest struct {
		UserID  uint  `json:"user_id"`
		Balance float64 `json:"balance"`
	}
	if err := c.ShouldBindJSON(&bodyRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	var existingUser model.User
	if err := initializers.DB.First(&existingUser, "id = ?", bodyRequest.UserID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Create a new customer
	customer := model.Customer{
		UserID:  bodyRequest.UserID,
		Balance: bodyRequest.Balance,
	}

	// Perform database operation to create the customer
	result := initializers.DB.Create(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Return the created customer
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}





