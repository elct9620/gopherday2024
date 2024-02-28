package usecase

import (
	"context"
	"time"

	"github.com/elct9620/gopherday2024/internal/entity"
)

type CreateShipmentCommandInput struct {
	ID string
}

type CreateShipmentCommandOutput struct {
	ID        string
	State     string
	UpdatedAt *time.Time
}

type CreateShipmentCommand struct {
	shipments ShipmentRepository
}

func NewCreateShipmentCommand(shipments ShipmentRepository) *CreateShipmentCommand {
	return &CreateShipmentCommand{
		shipments: shipments,
	}
}

func (c *CreateShipmentCommand) Execute(ctx context.Context, input *CreateShipmentCommandInput) (*CreateShipmentCommandOutput, error) {
	shipment := entity.NewShipment(input.ID)

	if err := c.shipments.Save(ctx, shipment); err != nil {
		return nil, err
	}

	return &CreateShipmentCommandOutput{
		ID:        shipment.ID,
		State:     string(shipment.State),
		UpdatedAt: shipment.UpdatedAt,
	}, nil
}
