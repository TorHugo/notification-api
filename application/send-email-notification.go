package application

import (
	"fmt"
	"notification-api/domain"
	"notification-api/infrastructure/config/mail"

	"gopkg.in/gomail.v2"
)

func SendEmailNotification(notification domain.Notification) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", mail.SMTP.Username)
	msg.SetHeader("To", notification.Contact)
	msg.SetHeader("Subject", notification.Subject)
	msg.SetBody("text/plain", notification.Body)

	dialer := mail.GetDialer()
	if err := dialer.DialAndSend(msg); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
