package event

import (
	"notification-api/application"
	"notification-api/infrastructure/config/database"
	"notification-api/infrastructure/config/repository"
)

func Start() *Publisher {
	eventRepository := repository.NewEventRepository(database.DB)
	createEvent := application.NewCreateEvent(eventRepository)
	eventPublisher := NewEventPublisher(createEvent)
	eventPublisher.Listen()
	return eventPublisher
}
