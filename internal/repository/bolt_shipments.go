package repository

import (
	"context"
	"encoding/json"
	"sort"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/elct9620/gopherday2024/internal/event"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"go.etcd.io/bbolt"
)

type BoltShipmentSchema struct {
	Type    string `json:"type"`
	Payload string `json:"payload"`
}

var _ usecase.ShipmentRepository = &BoltShipmentRepository{}

type BoltShipmentRepository struct {
	bucketName string
	db         *bbolt.DB
}

func NewBoltShipmentRepository(db *bbolt.DB) *BoltShipmentRepository {
	return &BoltShipmentRepository{
		bucketName: "shipments",
		db:         db,
	}
}

func (r *BoltShipmentRepository) Find(ctx context.Context, id string) (*entity.Shipment, error) {
	var events []event.ShipmentEvent = make([]event.ShipmentEvent, 0)

	err := r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.bucketName))
		if b == nil {
			return nil
		}

		shipmentBucket := b.Bucket([]byte(id))
		if shipmentBucket == nil {
			return usecase.ErrShipmentNotFound
		}

		return shipmentBucket.ForEach(func(k, v []byte) error {
			var value BoltShipmentSchema
			if err := json.Unmarshal(v, &value); err != nil {
				return err
			}

			switch value.Type {
			case "created":
				var payload event.ShipmentCreatedEvent
				if err := json.Unmarshal([]byte(value.Payload), &payload); err != nil {
					return err
				}

				events = append(events, &payload)
			case "item_added":
				var payload event.ShipmentItemAddedEvent
				if err := json.Unmarshal([]byte(value.Payload), &payload); err != nil {
					return err
				}
				events = append(events, &payload)
			case "state_shipping":
				var payload event.ShipmentShippingEvent
				if err := json.Unmarshal([]byte(value.Payload), &payload); err != nil {
					return err
				}
				events = append(events, &payload)
			case "state_delivered":
				var payload event.ShipmentDeliveredEvent
				if err := json.Unmarshal([]byte(value.Payload), &payload); err != nil {
					return err
				}
				events = append(events, &payload)
			}

			return nil
		})
	})

	if err != nil {
		return nil, err
	}

	if len(events) == 0 {
		return nil, usecase.ErrShipmentNotFound
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].CreatedAt().Before(events[j].CreatedAt())
	})

	shipment := entity.NewShipmentFromEvents(events)
	return shipment, nil
}

func (r *BoltShipmentRepository) Save(ctx context.Context, shipment *entity.Shipment) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(r.bucketName))
		if err != nil {
			return err
		}

		shipmentBucket, err := b.CreateBucketIfNotExists([]byte(shipment.ID))
		if err != nil {
			return err
		}

		events := shipment.PendingEvents()

		for _, e := range events {
			schema := BoltShipmentSchema{}
			switch e.(type) {
			case *event.ShipmentCreatedEvent:
				schema.Type = "created"
			case *event.ShipmentItemAddedEvent:
				schema.Type = "item_added"
			case *event.ShipmentShippingEvent:
				schema.Type = "state_shipping"
			case *event.ShipmentDeliveredEvent:
				schema.Type = "state_delivered"
			}

			payload, err := json.Marshal(e)
			if err != nil {
				return err
			}
			schema.Payload = string(payload)

			value, err := json.Marshal(schema)
			if err != nil {
				return err
			}

			if err := shipmentBucket.Put([]byte(e.ID()), value); err != nil {
				return err
			}
		}

		return nil
	})
}
