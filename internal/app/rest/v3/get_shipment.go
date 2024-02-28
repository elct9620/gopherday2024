package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetShipment struct {
}

func NewGetShipment() *GetShipment {
	return &GetShipment{}
}

func (e *GetShipment) Method() string {
	return "GET"
}

func (e *GetShipment) Path() string {
	return "/shipments/{id}"
}

func (e *GetShipment) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Header().Add("Content-Type", "application/json")

	res := Shipment{
		ID:        id,
		State:     "unknown",
		Items:     []ShipmentItem{},
		UpdatedAt: nil,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
