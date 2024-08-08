package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"notification-api/application"
	"notification-api/infrastructure/util"
	"time"

	"github.com/gin-gonic/gin"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/event"
)

type NotificationRequest struct {
	Contact    string            `json:"contact"`
	Subject    string            `json:"subject"`
	Template   string            `json:"template"`
	Parameters []model.Parameter `json:"parameters"`
}

type NotificationController struct {
	eventPublisher *event.Publisher
}

func NewNotificationController(publisher *event.Publisher) NotificationController {
	return NotificationController{eventPublisher: publisher}
}

func (p *NotificationController) SendNotification(ctx *gin.Context) {
	body, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not read request body"})
		return
	}

	payload := string(body)
	var req NotificationRequest

	if err := json.Unmarshal(body, &req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	templateWithParameters := util.ProcessTemplate(req.Template, req.Parameters)
	notification := model.Notification{
		To:      req.Contact,
		Subject: req.Subject,
		Body:    templateWithParameters,
	}

	go func() {
		if err := application.SendNotification(notification); err != nil {
			println("Error sending email:", err.Error())
		}
	}()

	eventMessage := model.Event{
		ID:        uuid.New().String(),
		Type:      "NotificationSent",
		Payload:   payload,
		Timestamp: time.Now(),
	}
	p.eventPublisher.Publish(eventMessage)
	ctx.JSON(http.StatusOK, gin.H{"status": "Notification sent successfully!"})
}
