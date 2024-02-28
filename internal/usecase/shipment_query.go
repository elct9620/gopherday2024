package usecase

import (
	"context"
	"errors"
	"time"
)

type ShipmentQueryInput struct {
	ID string
}

type ShipmentQueryOutput struct {
	ID       string
	State    string
	UpdateAt *time.Time
}

type ShipmentQuery struct {
	shipments ShipmentRepository
}

func NewShipmentQuery(shipments ShipmentRepository) *ShipmentQuery {
	return &ShipmentQuery{
		shipments: shipments,
	}
}

func (q *ShipmentQuery) Execute(ctx context.Context, input *ShipmentQueryInput) (*ShipmentQueryOutput, error) {
	shipment, err := q.shipments.Find(ctx, input.ID)
	if errors.Is(err, ErrShipmentNotFound) {
		return &ShipmentQueryOutput{
			ID:       input.ID,
			State:    "unknown",
			UpdateAt: nil,
		}, nil
	}

	// NOTE: we should not return the repository error directly
	if err != nil {
		return nil, err
	}

	return &ShipmentQueryOutput{
		ID:       shipment.ID,
		State:    string(shipment.State),
		UpdateAt: shipment.UpdatedAt,
	}, nil
}
