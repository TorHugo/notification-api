package controller

import (
	"net/http"
	"time"

	"notification-api/application"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/event"
	"notification-api/infrastructure/util"

	"github.com/gin-gonic/gin"
)

type NotificationRequest struct {
	Identifier string            `json:"identifier"`
	Contact    string            `json:"contact"`
	Subject    string            `json:"subject"`
	Template   string            `json:"template"`
	Parameters []model.Parameter `json:"parameters"`
	CreatedAt  time.Time         `json:"created_at"`
}

type NotificationController struct {
	eventPublisher *event.Publisher
}

func NewNotificationController(publisher *event.Publisher) NotificationController {
	return NotificationController{eventPublisher: publisher}
}

func (p *NotificationController) SendNotification(ctx *gin.Context) {
	var req NotificationRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	processedTemplate := util.ProcessTemplate(req.Template, req.Parameters)
	notification := model.Notification{
		To:      req.Contact,
		Subject: req.Subject,
		Body:    processedTemplate,
	}
	if err := application.SendNotification(notification); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	eventMessage := model.Event{
		ID:        req.Identifier,
		Type:      "NotificationSent",
		Payload:   map[string]interface{}{"to": req.Contact, "subject": "Notification Subject", "body": processedTemplate},
		Timestamp: time.Now().Unix(),
	}
	p.eventPublisher.Publish(eventMessage)

	ctx.JSON(http.StatusOK, gin.H{"status": "notification sent", "processed_template": processedTemplate})
}
