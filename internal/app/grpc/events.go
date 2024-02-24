package grpc

import (
	"context"

	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/elct9620/gopherday2024/pkg/events"
)

var _ events.EventsServer = &EventsServer{}

type EventsServer struct {
	events.EventsServer
	eventQuery *usecase.EventQuery
}

func NewEventsServer(
	eventQuery *usecase.EventQuery,
) *EventsServer {
	return &EventsServer{
		eventQuery: eventQuery,
	}
}

func (s *EventsServer) List(ctx context.Context, in *events.ListEventRequest) (*events.ListEventResponse, error) {
	output, err := s.eventQuery.Execute(ctx, &usecase.EventQueryInput{})
	if err != nil {
		return nil, err
	}

	resp := &events.ListEventResponse{
		Events: make([]*events.Event, 0, len(output.Events)),
	}

	for _, event := range output.Events {
		resp.Events = append(resp.Events, &events.Event{
			Id:        event.ID,
			CreatedAt: event.CreatedAt.String(),
		})
	}

	return resp, nil
}
