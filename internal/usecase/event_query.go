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
	events EventRepository
}

func NewEventQuery(events EventRepository) *EventQuery {
	return &EventQuery{
		events: events,
	}
}

func (q *EventQuery) Execute(ctx context.Context, input *EventQueryInput) (*EventQueryOutput, error) {
	events, err := q.events.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	outputEvents := make([]Event, 0, len(events))
	for _, event := range events {
		outputEvents = append(outputEvents, Event{
			ID: event.ID,
		})
	}

	return &EventQueryOutput{
		Events: outputEvents,
	}, nil
}
