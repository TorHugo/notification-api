package main

import (
	"github.com/gin-gonic/gin"
	"notification-api/domain/model"
	"notification-api/infrastructure/config/event"
)

func main() {
	server := gin.Default()
	eventPublisher := event.NewEventPublisher()
	eventPublisher.Listen(func(event model.Event) {
		// save-event-use-case
	})

	err := server.Run(":9000")
	if err != nil {
		return
	}
}
