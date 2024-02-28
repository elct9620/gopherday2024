package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type GetShipment struct {
	shipmentQuery *usecase.ShipmentQuery
}

func NewGetShipment(shipmentQuery *usecase.ShipmentQuery) *GetShipment {
	return &GetShipment{
		shipmentQuery: shipmentQuery,
	}
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

	shipment, err := e.shipmentQuery.Execute(r.Context(), &usecase.ShipmentQueryInput{
		ID: id,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to query shipment: %v", err), http.StatusInternalServerError)
		return
	}

	res := Shipment{
		ID:        shipment.ID,
		State:     shipment.State,
		Items:     []ShipmentItem{},
		UpdatedAt: shipment.UpdateAt,
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
