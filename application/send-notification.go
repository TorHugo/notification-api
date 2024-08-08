package application

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/mail"
)

func SendNotification(notification model.Notification) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.SMTP.Username)
	msg.SetHeader("To", notification.To)
	msg.SetHeader("Subject", notification.Subject)
	msg.SetBody("text/plain", notification.Body)

	dialer := mail.GetDialer()
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
