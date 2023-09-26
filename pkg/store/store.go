package store

import (
	"context"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/sirupsen/logrus"
)

func New(ctx context.Context, logger *logrus.Logger) (Store, error) {
	db, err := badger.Open(badger.DefaultOptions("/data/db").WithLogger(logger))
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping badgerDB: %w", err)
	}

	return &store{db: db}, nil
}

type store struct {
	db *badger.DB
}

func (s store) Update() {
	// TODO implement me
	panic("implement me")
}

func (s store) Get() {
	// TODO implement me
	panic("implement me")
}

func (s store) Set() {
	// TODO implement me
	panic("implement me")
}

func (s store) Delete() {
	// TODO implement me
	panic("implement me")
}

type Store interface {
	Update()
	Get()
	Set()
	Delete()
}
