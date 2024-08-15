package service

import (
	"fmt"
	"notification-api/application"
	"notification-api/domain/model"
	"notification-api/infrastructure/controller/models"
	"notification-api/infrastructure/util"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (ns *NotificationService) ProcessNotification(req models.NotificationDTO) (model.Notification, error) {
	if req.Contact == "" || req.Template == "" {
		return model.Notification{}, fmt.Errorf("missing required fields")
	}

	templateWithParameters := util.ProcessTemplate(req.Template, req.Parameters)
	notification := model.Notification{
		Contact: req.Contact,
		Subject: req.Subject,
		Body:    templateWithParameters,
	}

	return notification, nil
}
func (ns *NotificationService) SendEmailNotification(notification model.Notification) {
	go func() {
		err := application.SendEmailNotification(notification)
		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
		}
	}()
}
func (ns *NotificationService) SendSmsNotification(notification model.Notification) error {
	err := application.SendSmsNotification(notification)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
		return err
	}
	return nil
}
