package model

import "time"

type Notification struct {
	Identifier string    `json:"identifier"`
	To         string    `json:"to"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created-at"`
}
