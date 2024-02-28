package v3

import "time"

type ShipmentItem struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Shipment struct {
	ID        string         `json:"id"`
	State     string         `json:"state"`
	Items     []ShipmentItem `json:"items"`
	UpdatedAt *time.Time     `json:"updated_at"`
}
