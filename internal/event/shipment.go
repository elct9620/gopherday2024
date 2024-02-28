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

type ShipmentItemAddedEvent struct {
	event
	itemID string
	name   string
}

func NewShipmentItemAddedEvent(id, aggregateID, itemID, name string, createdAt time.Time) *ShipmentItemAddedEvent {
	return &ShipmentItemAddedEvent{
		event: event{
			id:          id,
			aggregateID: aggregateID,
			createdAt:   createdAt,
		},
		itemID: itemID,
		name:   name,
	}
}

func (e *ShipmentItemAddedEvent) ItemID() string {
	return e.itemID
}

func (e *ShipmentItemAddedEvent) Name() string {
	return e.name
}
