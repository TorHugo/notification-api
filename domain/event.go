package domain

import "time"

type Event struct {
	ID        string
	Type      string
	Payload   string
	Timestamp time.Time
}
