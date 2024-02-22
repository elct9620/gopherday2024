package v1

import (
	"fmt"
	"net/http"
)

type GetEvents struct {
}

func NewGetEvents() *GetEvents {
	return &GetEvents{}
}

func (e *GetEvents) Method() string {
	return "GET"
}

func (e *GetEvents) Path() string {
	return "/events"
}

func (e *GetEvents) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "[]")
}
