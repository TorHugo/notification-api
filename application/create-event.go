package application

import (
	"notification-api/domain"
	"notification-api/infrastructure/repository"
)

type CreateEvent struct {
	eventRepository *repository.EventRepository
}

func NewCreateEvent(repo *repository.EventRepository) *CreateEvent {
	return &CreateEvent{
		eventRepository: repo,
	}
}

func (ce *CreateEvent) Execute(event domain.Event) error {
	err := ce.eventRepository.Save(event)
	if err != nil {
		return err
	}
	return nil
}
