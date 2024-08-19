package service

import (
	"fmt"
	"notification-api/application"
	"notification-api/domain"
	"notification-api/infrastructure/controller/models"
	"notification-api/infrastructure/util"
)

type NotificationService struct {
}

func NewNotificationService() *NotificationService {
	return &NotificationService{}
}

func (ns *NotificationService) ProcessNotification(req models.NotificationDTO) (domain.Notification, error) {
	if req.Contact == "" || req.Template == "" {
		return domain.Notification{}, fmt.Errorf("missing required fields")
	}

	templateWithParameters := util.ProcessTemplate(req.Template, req.Parameters)
	notification := domain.Notification{
		Contact: req.Contact,
		Subject: req.Subject,
		Body:    templateWithParameters,
	}

	return notification, nil
}
func (ns *NotificationService) SendEmailNotification(notification domain.Notification) {
	go func() {
		err := application.SendEmailNotification(notification)
		if err != nil {
			fmt.Printf("Error sending notification: %v\n", err)
			// TODO: Send to retry process.
		}
	}()
}
func (ns *NotificationService) SendSmsNotification(notification domain.Notification) error {
	err := application.SendSmsNotification(notification)
	if err != nil {
		fmt.Printf("Error sending notification: %v\n", err)
		return err
	}
	return nil
}
