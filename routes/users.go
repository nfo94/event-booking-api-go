package routes

import (
	"event-booking-api-go/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.Signup()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func login(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login"})
}
