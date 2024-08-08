package main

import (
	"github.com/gin-gonic/gin"
	"notification-api/infrastructure/config/database"
	"notification-api/infrastructure/config/event"
	"notification-api/infrastructure/config/mail"
	"notification-api/infrastructure/controller"
)

func main() {

	mail.Start()
	database.Start()
	eventPublisher := event.Start()
	notificationController := controller.NewNotificationController(eventPublisher)

	r := gin.Default()
	r.POST("/api/send-notification", notificationController.SendNotification)

	r.Run(":8000")
}
