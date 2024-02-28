package event

import (
	"encoding/json"
	"time"
)

type Event interface {
	ID() string
	AggregateID() string
	CreatedAt() time.Time
}

var _ Event = &event{}
var _ json.Marshaler = &event{}
var _ json.Unmarshaler = &event{}

type event struct {
	id          string
	aggregateID string
	createdAt   time.Time
}

func (e *event) ID() string {
	return e.id
}

func (e *event) AggregateID() string {
	return e.aggregateID
}

func (e *event) CreatedAt() time.Time {
	return e.createdAt
}

func (e *event) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID          string    `json:"id"`
		AggregateID string    `json:"aggregate_id"`
		CreatedAt   time.Time `json:"created_at"`
	}{
		ID:          e.id,
		AggregateID: e.aggregateID,
		CreatedAt:   e.createdAt,
	})
}

func (e *event) UnmarshalJSON(data []byte) error {
	var v struct {
		ID          string    `json:"id"`
		AggregateID string    `json:"aggregate_id"`
		CreatedAt   time.Time `json:"created_at"`
	}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e.id = v.ID
	e.aggregateID = v.AggregateID
	e.createdAt = v.CreatedAt
	return nil
}
