package v3

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PostShipmentPayload struct {
	ID string `json:"id"`
}

type PostShipments struct {
}

func NewPostShipments() *PostShipments {
	return &PostShipments{}
}

func (e *PostShipments) Method() string {
	return "POST"
}

func (e *PostShipments) Path() string {
	return "/shipments"
}

func (e *PostShipments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var payload PostShipmentPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, fmt.Sprintf("failed to decode request: %v", err), http.StatusBadRequest)
		return
	}

	w.Header().Add("Content-Type", "application/json")

	createdAt := time.Now()
	shipment := Shipment{
		ID:        payload.ID,
		State:     "pending",
		Items:     []ShipmentItem{},
		UpdatedAt: &createdAt,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
