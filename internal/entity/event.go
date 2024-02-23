package entity

import "time"

type Event struct {
	ID        string
	CreatedAt time.Time
}

func NewEvent(id string) *Event {
	return &Event{
		ID:        id,
		CreatedAt: time.Now(),
	}
}
