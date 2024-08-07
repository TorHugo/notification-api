package event

import "notification-api/domain/model"

type Publisher struct {
	eventChannel chan model.Event
}

func NewEventPublisher() *Publisher {
	return &Publisher{
		eventChannel: make(chan model.Event, 100),
	}
}

func (ep *Publisher) Publish(event model.Event) {
	go func() {
		ep.eventChannel <- event
	}()
}
