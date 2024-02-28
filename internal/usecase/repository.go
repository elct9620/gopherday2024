package usecase

import (
	"context"
	"errors"

	"github.com/elct9620/gopherday2024/internal/entity"
)

var (
	// NOTE: usecase should not knows about the repository implementation
	ErrShipmentNotFound = errors.New("shipment not found")
)

type EventRepository interface {
	FindAll(context.Context) ([]*entity.Event, error)
	Save(context.Context, *entity.Event) error
}

type ShipmentRepository interface {
	Find(context.Context, string) (*entity.Shipment, error)
}
