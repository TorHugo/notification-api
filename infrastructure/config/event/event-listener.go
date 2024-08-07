package event

import "notification-api/domain/model"

func (ep *Publisher) Listen(handler func(event model.Event)) {
	go func() {
		for event := range ep.eventChannel {
			handler(event)
		}
	}()
}
