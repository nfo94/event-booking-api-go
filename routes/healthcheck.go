package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is running"})
}
