package usecase

import (
	"context"

	"github.com/elct9620/gopherday2024/internal/entity"
)

type EventRepository interface {
	FindAll(context.Context) ([]*entity.Event, error)
	Save(context.Context, *entity.Event) error
}
