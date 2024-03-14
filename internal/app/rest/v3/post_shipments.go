package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
)

type PostShipmentPayload struct {
	ID string `json:"id"`
}

type PostShipments struct {
	createShipmentCommand *usecase.CreateShipmentCommand
}

func NewPostShipments(
	createShipmentCommand *usecase.CreateShipmentCommand,
) *PostShipments {
	return &PostShipments{
		createShipmentCommand: createShipmentCommand,
	}
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

	output, err := e.createShipmentCommand.Execute(r.Context(), &usecase.CreateShipmentCommandInput{
		ID: payload.ID,
	})
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to create shipment: %v", err), http.StatusInternalServerError)
		return
	}

	shipment := Shipment{
		ID:        output.ID,
		State:     output.State,
		Items:     []ShipmentItem{},
		UpdatedAt: output.UpdatedAt,
	}

	if err := json.NewEncoder(w).Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
