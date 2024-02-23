package v1

import (
	"encoding/json"
	"net/http"
)

type PostEvents struct {
}

func NewPostEvents() *PostEvents {
	return &PostEvents{}
}

func (e *PostEvents) Method() string {
	return "POST"
}

func (e *PostEvents) Path() string {
	return "/events"
}

func (e *PostEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(json.RawMessage(`{"ok": true}`))
}
