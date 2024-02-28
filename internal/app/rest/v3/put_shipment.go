package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type PutShipmentPayload struct {
	State string `json:"state"`
}

type PutShipment struct {
	changeShipmentStateCommand *usecase.ChangeShipmentStateCommand
}

func NewPutShipment(
	changeShipmentStateCommand *usecase.ChangeShipmentStateCommand,
) *PutShipment {
	return &PutShipment{
		changeShipmentStateCommand: changeShipmentStateCommand,
	}
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

	output, err := e.changeShipmentStateCommand.Execute(r.Context(), &usecase.ChangeShipmentStateCommandInput{
		ShipmentID: id,
		State:      payload.State,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to change shipment state: %v", err), http.StatusInternalServerError)
		return
	}

	shipment := Shipment{
		ID:        output.ID,
		State:     output.State,
		UpdatedAt: output.UpdatedAt,
		Items:     []ShipmentItem{},
	}

	for _, item := range output.Items {
		shipment.Items = append(shipment.Items, ShipmentItem{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
