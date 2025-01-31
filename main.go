package main

import (
	"event-booking-api-go/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Run(":8080")
}
