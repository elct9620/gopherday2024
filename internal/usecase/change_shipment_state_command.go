package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/elct9620/gopherday2024/internal/entity"
)

type ChangeShipmentStateCommandInput struct {
	ShipmentID string
	State      string
}

type ChangeShipmentStateCommandOutput struct {
	ID    string
	State string
	Items []struct {
		ID   string
		Name string
	}
	UpdatedAt *time.Time
}

type ChangeShipmentStateCommand struct {
	shipments ShipmentRepository
}

func NewChangeShipmentStateCommand(shipments ShipmentRepository) *ChangeShipmentStateCommand {
	return &ChangeShipmentStateCommand{
		shipments: shipments,
	}
}

func (c *ChangeShipmentStateCommand) Execute(ctx context.Context, input *ChangeShipmentStateCommandInput) (*ChangeShipmentStateCommandOutput, error) {
	shipment, err := c.shipments.Find(ctx, input.ShipmentID)
	if err != nil {
		return nil, err
	}

	switch entity.ShipmentState(input.State) {
	case entity.ShipmentStateShipping:
		shipment.ShipAt(time.Now())
	case entity.ShipmentStateDelivered:
		shipment.DeliveredAt(time.Now())
	default:
		return nil, fmt.Errorf("invalid state: %s", input.State)
	}

	if err := c.shipments.Save(ctx, shipment); err != nil {
		return nil, err
	}

	output := &ChangeShipmentStateCommandOutput{
		ID:        shipment.ID,
		State:     string(shipment.State),
		UpdatedAt: shipment.UpdatedAt,
	}

	for _, item := range shipment.Items {
		output.Items = append(output.Items, struct {
			ID   string
			Name string
		}{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	return output, nil
}
