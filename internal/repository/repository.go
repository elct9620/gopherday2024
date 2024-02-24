package repository

import (
	"github.com/elct9620/gopherday2024/internal/usecase"
	"github.com/google/wire"
)

var DefaultSet = wire.NewSet(
	wire.Bind(new(usecase.EventRepository), new(*BoltEventRepository)),
	NewBoltEventRepository,
)

var InMemorySet = wire.NewSet(
	wire.Bind(new(usecase.EventRepository), new(*InMemoryEventRepository)),
	NewInMemoryEventRepository,
)
