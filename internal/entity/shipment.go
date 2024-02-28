package entity

import (
	"sync"
	"time"

	"github.com/elct9620/gopherday2024/internal/event"
	"github.com/google/uuid"
)

type ShipmentState string

const (
	ShipmentStatePending   ShipmentState = "pending"
	ShipmentStateShipping  ShipmentState = "shipping"
	ShipmentStateDelivered ShipmentState = "delivered"
)

type ShipmentItem struct {
	ID   string
	Name string
}

type Shipment struct {
	ID        string
	State     ShipmentState
	UpdatedAt *time.Time
	Items     []ShipmentItem

	// NOTE: to prevent event appened when we are reading the pending events
	mux           sync.RWMutex
	pendingEvents []event.ShipmentEvent
}

func NewShipment(id string) *Shipment {
	shipment := &Shipment{}

	event := event.NewShipmentCreatedEvent(uuid.NewString(), id, time.Now())
	shipment.pendingEvents = append(shipment.pendingEvents, event)
	shipment.apply(event)

	return shipment
}

func NewShipmentFromEvents(events []event.ShipmentEvent) *Shipment {
	shipment := &Shipment{}

	for _, e := range events {
		shipment.apply(e)
	}

	return shipment
}

func (s *Shipment) apply(e event.ShipmentEvent) {
	switch e := e.(type) {
	case *event.ShipmentCreatedEvent:
		createdAt := e.CreatedAt()
		s.ID = e.AggregateID()
		s.State = ShipmentStatePending
		s.UpdatedAt = &createdAt
	case *event.ShipmentItemAddedEvent:
		s.Items = append(s.Items, ShipmentItem{
			ID:   e.ItemID(),
			Name: e.Name(),
		})
	}
}

func (s *Shipment) AddItem(id, name string) {
	event := event.NewShipmentItemAddedEvent(uuid.NewString(), s.ID, id, name, time.Now())
	s.pendingEvents = append(s.pendingEvents, event)
	s.apply(event)
}

func (s *Shipment) PendingEvents() []event.ShipmentEvent {
	s.mux.RLock()
	defer s.mux.RUnlock()

	events := make([]event.ShipmentEvent, len(s.pendingEvents))
	copy(events, s.pendingEvents)

	return events
}

func (s *Shipment) ClearPendingEvents() {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.pendingEvents = make([]event.ShipmentEvent, 0)
}
