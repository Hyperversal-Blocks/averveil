package store

import (
	"fmt"

	"github.com/dgraph-io/badger/v4"
)

func (s *store) Get(key string) ([]byte, error) {
	// Retrieve the data from BadgerDB
	var obj []byte
	err := s.badgerDB.View(func(txn *badger.Txn) error {
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
	s.badgerMutex.Lock()
	defer s.badgerMutex.Unlock()
	err := s.badgerDB.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	})
	if err != nil {
		return fmt.Errorf("unable to create object: %w", err)
	}
	return nil
}

func (s *store) Delete(key string) error {
	s.badgerMutex.Lock()
	defer s.badgerMutex.Unlock()
	// Start a writable transaction.
	err := s.badgerDB.Update(func(txn *badger.Txn) error {
		// Delete the key-value pair associated with the specified key
		return txn.Delete([]byte(key))
	})

	if err != nil {
		return fmt.Errorf("unable to delete a key and corresponding val: %w", err)
	}

	return nil
}
