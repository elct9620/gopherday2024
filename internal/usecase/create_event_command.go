package usecase

import "context"

type CreateEventCommandInput struct {
}

type CreateEventCommandOutput struct {
	ID string
}

type CreateEventCommand struct {
}

func NewCreateEventCommand() *CreateEventCommand {
	return &CreateEventCommand{}
}

func (c *CreateEventCommand) Execute(ctx context.Context, input *CreateEventCommandInput) (*CreateEventCommandOutput, error) {
	return &CreateEventCommandOutput{
		ID: "645fc1c8-3505-42b5-913d-bf2a84fd70f1",
	}, nil
}
