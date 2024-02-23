package usecase

import "context"

type EventQueryInput struct {
}

type Event struct {
	ID string
}

type EventQueryOutput struct {
	Events []Event
}

type EventQuery struct {
}

func NewEventQuery() *EventQuery {
	return &EventQuery{}
}

func (q *EventQuery) Execute(ctx context.Context, input *EventQueryInput) (*EventQueryOutput, error) {
	return &EventQueryOutput{
		Events: []Event{},
	}, nil
}
