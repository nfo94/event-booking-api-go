package routes

import (
	"event-booking-api-go/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func createEventRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	event, err := models.GetEventbyId(eventId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event ID"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User successfully register for event"})
}

func cancelEventRegistration(c *gin.Context) {
	userId := c.GetInt64("userId")

	eventId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event ID"})
		return
	}

	var event models.Event
	event.ID = eventId

	err = event.Cancel(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register event"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Event registration successfully cancelled"})
}
