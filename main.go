package main

import (
	"net/http"

	"event-booking-api-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/healthcheck", healthcheck)
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080")
}

func healthcheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "API is running"})
}

func getEvents(c *gin.Context) {
	events := models.GetAllEvents()
	c.JSON(http.StatusOK, events)
}

func createEvent(c *gin.Context) {
	var event models.Event
	err := c.ShouldBindJSON(&event)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Save()
	c.JSON(http.StatusCreated, gin.H{"message": "Event created"})
}
