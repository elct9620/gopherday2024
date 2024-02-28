package v3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type PostShipmentItemPayload struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type PostShipmentItems struct {
	createShipmentCommand *usecase.CreateShipmentCommand
}

func NewPostShipmentItems() *PostShipmentItems {
	return &PostShipmentItems{}
}

func (e *PostShipmentItems) Method() string {
	return "POST"
}

func (e *PostShipmentItems) Path() string {
	return "/shipments/{id}/items"
}

func (e *PostShipmentItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload PostShipmentItemPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	nowTime := time.Now()
	shipment := Shipment{
		ID:    id,
		State: "pending",
		Items: []ShipmentItem{
			{
				ID:   payload.ID,
				Name: payload.Name,
			},
		},
		UpdatedAt: &nowTime,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
