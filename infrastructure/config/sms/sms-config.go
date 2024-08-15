package sms

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type TwilioConfig struct {
	AccountSID string
	AuthToken  string
	FromPhone  string
}

var Twilio *TwilioConfig

func Start() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	Twilio = &TwilioConfig{
		AccountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
		AuthToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		FromPhone:  os.Getenv("TWILIO_FROM_PHONE"),
	}
}
