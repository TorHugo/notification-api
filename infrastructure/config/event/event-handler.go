package event

import (
	"notification-api/application"
	"notification-api/domain"
)

type Publisher struct {
	eventChannel chan domain.Event
	createEvent  *application.CreateEvent
}

func NewEventPublisher(useCase *application.CreateEvent) *Publisher {
	return &Publisher{
		eventChannel: make(chan domain.Event, 100),
		createEvent:  useCase,
	}
}

func (ep *Publisher) Publish(event domain.Event) {
	go func() {
		ep.eventChannel <- event
	}()
}

func (ep *Publisher) Listen() {
	go func() {
		for event := range ep.eventChannel {
			err := ep.createEvent.Execute(event)
			if err != nil {
				println("Error saving event:", err.Error())
			}
		}
	}()
}
