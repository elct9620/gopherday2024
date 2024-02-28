package event

import "time"

type Event interface {
	ID() string
	AggregateID() string
	CreatedAt() time.Time
}

var _ Event = &event{}

type event struct {
	id          string
	aggregateID string
	createdAt   time.Time
}

func (e *event) ID() string {
	return e.id
}

func (e *event) AggregateID() string {
	return e.aggregateID
}

func (e *event) CreatedAt() time.Time {
	return e.createdAt
}
