package repository

import (
	"context"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/elct9620/gopherday2024/internal/usecase"
)

var _ usecase.EventRepository = &InMemoryEventRepository{}

type InMemoryEventRepository struct {
	entities []*entity.Event
}

func NewInMemoryEventRepository() *InMemoryEventRepository {
	return &InMemoryEventRepository{
		entities: []*entity.Event{},
	}
}

func (r *InMemoryEventRepository) FindAll(ctx context.Context) ([]*entity.Event, error) {
	return r.entities, nil
}
