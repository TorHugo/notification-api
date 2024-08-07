package application

import "notification-api/domain/model"

type CreateEvent struct {
	// gateway
}

func NewCreateEvent() CreateEvent {
	return CreateEvent{}
}

func (createEvent *CreateEvent) Execute(model.Event, error) {

}
