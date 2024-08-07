package application

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/mail"
)

type SmtpClient struct {
	host     string
	port     int
	username string
	password string
}

func execute(notification model.Notification) error {
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
