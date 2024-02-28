package usecase

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type CreateShipmentItemCommandInput struct {
	ShipmentID string
	Name       string
}

type CreateShipmentItemCommandOutput struct {
	ID    string
	State string
	Items []struct {
		ID   string
		Name string
	}
	UpdatedAt *time.Time
}

type CreateShipmentItemCommand struct {
	shipments ShipmentRepository
}

func NewCreateShipmentItemCommand(shipments ShipmentRepository) *CreateShipmentItemCommand {
	return &CreateShipmentItemCommand{
		shipments: shipments,
	}
}

func (c *CreateShipmentItemCommand) Execute(ctx context.Context, input *CreateShipmentItemCommandInput) (*CreateShipmentItemCommandOutput, error) {
	shipment, err := c.shipments.Find(ctx, input.ShipmentID)
	if err != nil {
		return nil, err
	}

	shipment.AddItem(uuid.NewString(), input.Name)

	if err := c.shipments.Save(ctx, shipment); err != nil {
		return nil, err
	}

	output := &CreateShipmentItemCommandOutput{
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
