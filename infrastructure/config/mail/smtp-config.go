package mail

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Dialer   *gomail.Dialer
}

var SMTP *SMTPClient

var host = os.Getenv("SMTP_HOST")
var username = os.Getenv("SMTP_USERNAME")
var password = os.Getenv("SMTP_PASSWORD")

func Start() {
	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		fmt.Println("Error converting string to int:", err)
		return
	}
	SMTP = &SMTPClient{
		Dialer: gomail.NewDialer(host, port, username, password),
	}
}

func GetDialer() *gomail.Dialer {
	return SMTP.Dialer
}
