package routes

import (
	"github.com/gin-gonic/gin"
	"notification-api/infrastructure/config/event"
	"notification-api/infrastructure/controller"
	"notification-api/infrastructure/service"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	eventPublisher := event.Start()
	notificationService := service.NewNotificationService()
	notificationController := controller.NewNotificationController(notificationService, eventPublisher)

	r.POST("/api/mail/send-notification", notificationController.SendEmailNotification)
	r.POST("/api/sms/send-notification", notificationController.SendSmsNotification)
	return r
}
