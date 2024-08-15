package application

import (
	"fmt"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/sms"
)

func SendSmsNotification(notification model.Notification) error {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: sms.Twilio.AccountSID,
		Password: sms.Twilio.AuthToken,
	})

	params := &openapi.CreateMessageParams{}
	params.SetTo(notification.Contact)
	params.SetFrom(sms.Twilio.FromPhone)
	params.SetBody(notification.Body)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("failed to send sms: %w", err)
	}
	return nil
}
