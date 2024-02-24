package app

import (
	"time"

	"go.etcd.io/bbolt"
)

func ProvideBoltDB() (*bbolt.DB, func(), error) {
	db, err := bbolt.Open("db/bolt.db", 0600, &bbolt.Options{
		Timeout: 10 * time.Second,
	})

	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		db.Close()
	}

	return db, cleanup, nil
}
