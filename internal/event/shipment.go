package event

import (
	"encoding/json"
	"time"
)

type ShipmentEvent interface {
	Event
}

var _ ShipmentEvent = &ShipmentCreatedEvent{}
var _ json.Marshaler = &ShipmentCreatedEvent{}
var _ json.Unmarshaler = &ShipmentCreatedEvent{}

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

var _ ShipmentEvent = &ShipmentItemAddedEvent{}
var _ json.Marshaler = &ShipmentItemAddedEvent{}
var _ json.Unmarshaler = &ShipmentItemAddedEvent{}

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

func (e *ShipmentItemAddedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          string    `json:"id"`
		AggregateID string    `json:"aggregate_id"`
		CreatedAt   time.Time `json:"created_at"`
		ItemID      string    `json:"item_id"`
		Name        string    `json:"name"`
	}{
		ID:          e.id,
		AggregateID: e.aggregateID,
		CreatedAt:   e.createdAt,
		ItemID:      e.itemID,
		Name:        e.name,
	})
}

func (e *ShipmentItemAddedEvent) UnmarshalJSON(data []byte) error {
	var v struct {
		ID          string    `json:"id"`
		AggregateID string    `json:"aggregate_id"`
		CreatedAt   time.Time `json:"created_at"`
		ItemID      string    `json:"item_id"`
		Name        string    `json:"name"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e.id = v.ID
	e.aggregateID = v.AggregateID
	e.createdAt = v.CreatedAt
	e.itemID = v.ItemID
	e.name = v.Name
	return nil
}

var _ ShipmentEvent = &ShipmentShippingEvent{}
var _ json.Marshaler = &ShipmentShippingEvent{}
var _ json.Unmarshaler = &ShipmentShippingEvent{}

type ShipmentShippingEvent struct {
	event
}

func NewShipmentShippingEvent(id, aggregateID string, createdAt time.Time) *ShipmentShippingEvent {
	return &ShipmentShippingEvent{
		event: event{
			id:          id,
			aggregateID: aggregateID,
			createdAt:   createdAt,
		},
	}
}

var _ ShipmentEvent = &ShipmentDeliveredEvent{}
var _ json.Marshaler = &ShipmentDeliveredEvent{}
var _ json.Unmarshaler = &ShipmentDeliveredEvent{}

type ShipmentDeliveredEvent struct {
	event
}

func NewShipmentDeliveredEvent(id, aggregateID string, createdAt time.Time) *ShipmentDeliveredEvent {
	return &ShipmentDeliveredEvent{
		event: event{
			id:          id,
			aggregateID: aggregateID,
			createdAt:   createdAt,
		},
	}
}
