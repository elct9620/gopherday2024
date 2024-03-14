package v3

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/go-chi/chi/v5"
)

type PostShipmentItemPayload struct {
	Name string `json:"name"`
}

type PostShipmentItems struct {
	createShipmentItemCommand *usecase.CreateShipmentItemCommand
}

func NewPostShipmentItems(
	createShipmentItemCommand *usecase.CreateShipmentItemCommand,
) *PostShipmentItems {
	return &PostShipmentItems{
		createShipmentItemCommand: createShipmentItemCommand,
	}
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

	output, err := e.createShipmentItemCommand.Execute(r.Context(), &usecase.CreateShipmentItemCommandInput{
		ShipmentID: id,
		Name:       payload.Name,
	})

	if err != nil {
		http.Error(w, fmt.Sprintf("failed to execute command: %v", err), http.StatusInternalServerError)
		return
	}

	shipment := Shipment{
		ID:        output.ID,
		State:     output.State,
		UpdatedAt: output.UpdatedAt,
		Items:     make([]ShipmentItem, 0, len(output.Items)),
	}

	for _, item := range output.Items {
		shipment.Items = append(shipment.Items, ShipmentItem{
			ID:   item.ID,
			Name: item.Name,
		})
	}

	if err := json.NewEncoder(w).Encode(shipment); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
	}
}
