package usecase

import (
	"context"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/google/uuid"
)

type CreateEventCommandInput struct {
}

type CreateEventCommandOutput struct {
	ID string
}

type CreateEventCommand struct {
	events EventRepository
}

func NewCreateEventCommand(events EventRepository) *CreateEventCommand {
	return &CreateEventCommand{
		events: events,
	}
}

func (c *CreateEventCommand) Execute(ctx context.Context, input *CreateEventCommandInput) (*CreateEventCommandOutput, error) {
	event := entity.NewEvent(uuid.NewString())

	if err := c.events.Save(ctx, event); err != nil {
		return nil, err
	}

	return &CreateEventCommandOutput{
		ID: event.ID,
	}, nil
}
