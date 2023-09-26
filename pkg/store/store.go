package store

import (
	"context"
	"fmt"
	"sync"

	"github.com/dgraph-io/badger/v4"
	"github.com/sirupsen/logrus"
)

func New(ctx context.Context, logger *logrus.Logger) (Store, error) {
	db, err := badger.Open(badger.DefaultOptions("/data/db").WithLogger(logger))
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping badgerDB: %w", err)
	}
	return &store{db: db, m: &sync.Mutex{}}, nil
}

type store struct {
	db *badger.DB
	m  *sync.Mutex
}

func (s *store) Get(key string) ([]byte, error) {
	// Retrieve the data from BadgerDB
	var obj []byte
	err := s.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		obj, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch user object: %w", err)
	}

	return obj, nil
}

func (s *store) Put(key string, value []byte) error {
	s.m.Lock()
	defer s.m.Unlock()
	err := s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
	if err != nil {
		return fmt.Errorf("unable to create object: %w", err)
	}
	return nil
}

func (s *store) Delete(key string) error {
	s.m.Lock()
	defer s.m.Unlock()
	// Start a writable transaction.
	err := s.db.Update(func(txn *badger.Txn) error {
		// Delete the key-value pair associated with the specified key
		return txn.Delete([]byte(key))
	})

	if err != nil {
		return fmt.Errorf("unable to delete a key and corresponding val: %w", err)
	}

	return nil
}

type Store interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
}
