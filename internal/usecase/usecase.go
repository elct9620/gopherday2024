package usecase

import "github.com/google/wire"

var DefaultSet = wire.NewSet(
	NewEventQuery,
	NewCreateEventCommand,
	NewShipmentQuery,
	NewCreateShipmentCommand,
)
