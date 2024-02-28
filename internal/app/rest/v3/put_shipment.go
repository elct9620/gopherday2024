package v3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type PutShipmentPayload struct {
	State string `json:"state"`
}

type PutShipment struct {
}

func NewPutShipment() *PutShipment {
	return &PutShipment{}
}

func (e *PutShipment) Method() string {
	return "PUT"
}

func (e *PutShipment) Path() string {
	return "/shipments/{id}"
}

func (e *PutShipment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var payload PutShipmentPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	updatedAt := time.Now()
	shipment := Shipment{
		ID:        id,
		State:     payload.State,
		Items:     []ShipmentItem{},
		UpdatedAt: &updatedAt,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
