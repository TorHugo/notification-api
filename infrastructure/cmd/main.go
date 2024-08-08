package main

import (
	"github.com/gin-gonic/gin"
	"notification-api/infrastructure/config/database"
	"notification-api/infrastructure/config/event"
	"notification-api/infrastructure/config/mail"
	"notification-api/infrastructure/config/repository"
	"notification-api/infrastructure/controller"
)

func main() {

	mail.Start()
	database.Start()

	eventRepo := repository.NewEventRepository(database.DB)
	eventPublisher := event.NewEventPublisher(eventRepo)
	eventPublisher.Listen()

	notificationController := controller.NewNotificationController(eventPublisher)

	r := gin.Default()
	r.POST("/send-notification", notificationController.SendNotification)

	r.Run(":8080")
}
