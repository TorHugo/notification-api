package main

import (
	"notification-api/infrastructure/config/database"
	"notification-api/infrastructure/config/mail"
	"notification-api/infrastructure/config/sms"
	"notification-api/infrastructure/routes"
)

func main() {

	mail.Start()
	sms.Start()
	database.Start()
	r := routes.SetupRouter()
	err := r.Run(":8000")
	if err != nil {
		return
	}
}
