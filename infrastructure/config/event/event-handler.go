package event

import (
	"notification-api/domain/model"
	"notification-api/infrastructure/config/repository"
)

type Publisher struct {
	eventChannel chan model.Event
	repository   *repository.EventRepository
}

func NewEventPublisher(repo *repository.EventRepository) *Publisher {
	return &Publisher{
		eventChannel: make(chan model.Event, 100),
		repository:   repo,
	}
}

func (ep *Publisher) Publish(event model.Event) {
	go func() {
		ep.eventChannel <- event
	}()
}

func (ep *Publisher) Listen() {
	go func() {
		for event := range ep.eventChannel {
			err := ep.repository.Save(event)
			if err != nil {
				println("Error saving event:", err.Error())
			}
		}
	}()
}
