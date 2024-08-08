package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Dialer   *gomail.Dialer
}

var SMTP *SMTPClient

func Start() {
	errLoad := godotenv.Load()
	if errLoad != nil {
		log.Fatalf("Error loading .env file")
	}

	var host = os.Getenv("SMTP_HOST")
	var username = os.Getenv("SMTP_USERNAME")
	var password = os.Getenv("SMTP_PASSWORD")
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	SMTP = &SMTPClient{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Dialer:   gomail.NewDialer(host, port, username, password),
	}
}

func GetDialer() *gomail.Dialer {
	return SMTP.Dialer
}
