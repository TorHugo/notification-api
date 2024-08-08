package repository

import (
	"database/sql"
	"encoding/json"
	"notification-api/domain/model"
)

type EventRepository struct {
	db *sql.DB
}

func NewEventRepository(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) Save(event model.Event) error {
	payloadJSON, err := json.Marshal(event.Payload)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(`
        INSERT INTO events_tb (identifier, event_type, transaction, created_at)
        VALUES ($1, $2, $3, $4)
    `, event.ID, event.Type, payloadJSON, event.Timestamp)
	return err
}
