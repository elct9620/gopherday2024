package event

import "time"

type ShipmentEvent interface {
	Event
}

type ShipmentCreatedEvent struct {
	event
}

func NewShipmentCreatedEvent(id, aggregateID string, createdAt time.Time) *ShipmentCreatedEvent {
	return &ShipmentCreatedEvent{
		event: event{
			id:          id,
			aggregateID: aggregateID,
			createdAt:   createdAt,
		},
	}
}
