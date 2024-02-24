package repository

import (
	"context"
	"encoding/json"
	"time"

	"github.com/elct9620/gopherday2024/internal/entity"
	"github.com/elct9620/gopherday2024/internal/usecase"
	"go.etcd.io/bbolt"
)

type BoltEventSchema struct {
	CreatedAt time.Time `json:"created_at"`
}

var _ usecase.EventRepository = &BoltEventRepository{}

type BoltEventRepository struct {
	bucketName string
	db         *bbolt.DB
}

func NewBoltEventRepository(db *bbolt.DB) *BoltEventRepository {
	return &BoltEventRepository{
		bucketName: "events",
		db:         db,
	}
}

func (r *BoltEventRepository) FindAll(ctx context.Context) ([]*entity.Event, error) {
	entities := make([]*entity.Event, 0)

	r.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(r.bucketName))
		if b == nil {
			return nil
		}

		b.ForEach(func(k, v []byte) error {
			var value BoltEventSchema
			if err := json.Unmarshal(v, &value); err != nil {
				return err
			}

			entities = append(entities, &entity.Event{
				ID:        string(k),
				CreatedAt: value.CreatedAt,
			})

			return nil
		})

		return nil
	})

	return entities, nil
}

func (r *BoltEventRepository) Save(ctx context.Context, event *entity.Event) error {
	return r.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(r.bucketName))
		if err != nil {
			return err
		}

		value, err := json.Marshal(BoltEventSchema{
			CreatedAt: event.CreatedAt,
		})
		if err != nil {
			return err
		}

		return b.Put([]byte(event.ID), value)
	})
}
