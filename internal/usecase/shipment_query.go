package usecase

import (
	"context"
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
}

func NewShipmentQuery() *ShipmentQuery {
	return &ShipmentQuery{}
}

func (q *ShipmentQuery) Execute(ctx context.Context, input *ShipmentQueryInput) (*ShipmentQueryOutput, error) {
	return &ShipmentQueryOutput{
		ID:       input.ID,
		State:    "unknown",
		UpdateAt: nil,
	}, nil
}
