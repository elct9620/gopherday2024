package v2

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
)

type GetEvents struct {
	eventQuery *usecase.EventQuery
}

func NewGetEvents(
	eventQuery *usecase.EventQuery,
) *GetEvents {
	return &GetEvents{
		eventQuery: eventQuery,
	}
}

func (e *GetEvents) Method() string {
	return "GET"
}

func (e *GetEvents) Path() string {
	return "/events"
}

func (e *GetEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output, err := e.eventQuery.Execute(r.Context(), &usecase.EventQueryInput{})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	events := make([]Event, 0, len(output.Events))
	for _, event := range output.Events {
		events = append(events, Event{
			ID:        event.ID,
			CreatedAt: event.CreatedAt,
		})
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(events); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
