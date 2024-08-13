package model

import (
	"encoding/json"
	"time"
)

type Notification struct {
	Identifier string    `json:"identifier"`
	To         string    `json:"to"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
	CreatedAt  time.Time `json:"created-at"`
}

func (n *Notification) ToJSON() string {
	data, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(data)
}
