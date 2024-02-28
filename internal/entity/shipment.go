package entity

import "time"

type ShipmentState string

const (
	ShipmentStatePending   ShipmentState = "pending"
	ShipmentStateShipping  ShipmentState = "shipping"
	ShipmentStateDelivered ShipmentState = "delivered"
)

type Shipment struct {
	ID        string
	State     ShipmentState
	UpdatedAt *time.Time
}

func NewShipment(id string) *Shipment {
	return &Shipment{
		ID:    id,
		State: ShipmentStatePending,
	}
}
