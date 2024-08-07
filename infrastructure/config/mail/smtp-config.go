package mail

import "gopkg.in/gomail.v2"

type SMTPClient struct {
	Host     string
	Port     int
	Username string
	Password string
	Dialer   *gomail.Dialer
}

var SMTP *SMTPClient

func StartSmpt(host string, port int, username, password string) {
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
