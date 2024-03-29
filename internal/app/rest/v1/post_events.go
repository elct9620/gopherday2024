package v1

import (
	"encoding/json"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
)

type PostEventResponse struct {
	ID string `json:"id"`
}

type PostEvents struct {
	createEvent *usecase.CreateEventCommand
}

func NewPostEvents(
	createEvent *usecase.CreateEventCommand,
) *PostEvents {
	return &PostEvents{
		createEvent: createEvent,
	}
}

func (e *PostEvents) Method() string {
	return "POST"
}

func (e *PostEvents) Path() string {
	return "/events"
}

func (e *PostEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, err := e.createEvent.Execute(r.Context(), &usecase.CreateEventCommandInput{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	isSuccess := len(res.ID) > 0
	if !isSuccess {
		_, err = w.Write(json.RawMessage(`{"ok": false}`))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}

	response := PostEventResponse{
		ID: res.ID,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
