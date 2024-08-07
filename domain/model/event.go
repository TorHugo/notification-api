package model

type Event struct {
	ID        string
	Type      string
	Payload   map[string]interface{}
	Timestamp int64
}
