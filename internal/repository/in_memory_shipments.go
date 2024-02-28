package repository

import (
	"context"
	"sync"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/elct9620/gopherday2024/internal/event"
	"github.com/elct9620/gopherday2024/internal/usecase"
)

var _ usecase.ShipmentRepository = &InMemoryShipmentRepository{}

type InMemoryShipmentRepository struct {
	mux    sync.RWMutex
	events []event.ShipmentEvent
}

func NewInMemoryShipmentRepository() *InMemoryShipmentRepository {
	return &InMemoryShipmentRepository{
		events: make([]event.ShipmentEvent, 0),
	}
}

func (r *InMemoryShipmentRepository) Find(ctx context.Context, id string) (*entity.Shipment, error) {
	r.mux.RLock()
	defer r.mux.RUnlock()

	events := make([]event.ShipmentEvent, 0)
	for _, e := range r.events {
		if e.AggregateID() == id {
			events = append(events, e)
		}
	}

	if len(events) == 0 {
		return nil, usecase.ErrShipmentNotFound
	}
	shipment := entity.NewShipmentFromEvents(events)

	return shipment, nil
}

func (r *InMemoryShipmentRepository) Save(ctx context.Context, shipment *entity.Shipment) error {
	r.mux.Lock()
	defer r.mux.Unlock()

	events := shipment.PendingEvents()
	r.events = append(r.events, events...)
	shipment.ClearPendingEvents()

	return nil
}
