package repository

import (
	"context"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/elct9620/gopherday2024/internal/usecase"
)

var _ usecase.ShipmentRepository = &InMemoryShipmentRepository{}

type InMemoryShipmentRepository struct {
}

func NewInMemoryShipmentRepository() *InMemoryShipmentRepository {
	return &InMemoryShipmentRepository{}
}

func (r *InMemoryShipmentRepository) Find(ctx context.Context, id string) (*entity.Shipment, error) {
	return nil, usecase.ErrShipmentNotFound
}

func (r *InMemoryShipmentRepository) Save(ctx context.Context, shipment *entity.Shipment) error {
	return nil
}
