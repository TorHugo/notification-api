package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/event"
	"notification-api/infrastructure/controller/models"
	"notification-api/infrastructure/service"
	"time"
)

type NotificationController struct {
	service        *service.NotificationService
	eventPublisher *event.Publisher
}

func NewNotificationController(svc *service.NotificationService, publisher *event.Publisher) *NotificationController {
	return &NotificationController{service: svc, eventPublisher: publisher}
}
func (p *NotificationController) SendEmailNotification(ctx *gin.Context) {
	var req models.NotificationDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ApiResponse{Error: "Invalid input", Data: err.Error()})
		return
	}

	notification, err := p.service.ProcessNotification(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ApiResponse{Error: "Failed to process notification"})
		return
	}

	p.service.SendEmailNotification(notification)

	eventMessage := model.Event{
		ID:        uuid.New().String(),
		Type:      "SEND_EMAIL_NOTIFICATION",
		Payload:   notification.ToJSON(),
		Timestamp: time.Now(),
	}
	go p.eventPublisher.Publish(eventMessage)
	ctx.JSON(http.StatusOK, models.ApiResponse{Message: "Notification sent successfully!"})
}
func (p *NotificationController) SendSmsNotification(ctx *gin.Context) {
	var req models.NotificationDTO
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ApiResponse{Error: "Invalid input", Data: err.Error()})
		return
	}

	notification, err := p.service.ProcessNotification(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ApiResponse{Error: "Failed to process notification"})
		return
	}

	p.service.SendSmsNotification(notification)

	eventMessage := model.Event{
		ID:        uuid.New().String(),
		Type:      "SEND_SMS_NOTIFICATION",
		Payload:   notification.ToJSON(),
		Timestamp: time.Now(),
	}
	go p.eventPublisher.Publish(eventMessage)
	ctx.JSON(http.StatusOK, models.ApiResponse{Message: "Notification sent successfully!"})
}
