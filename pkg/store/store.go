package store

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"github.com/dgraph-io/badger/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func New(logger *logrus.Logger,
	path string, inMem, logging bool) (Store, error) {
	badgerDB, err := initBadger(logger, path, inMem, logging)
	if err != nil {
		return nil, err
	}

	sqlDB, err := initSQLite()
	if err != nil {
		return nil, err
	}

	return &store{
		badgerDB:    badgerDB,
		badgerMutex: &sync.Mutex{},
		sqlDB:       sqlDB,
		sqlMutex:    &sync.Mutex{},
	}, nil
}

func initBadger(logger *logrus.Logger,
	path string, inMem, logging bool) (*badger.DB, error) {
	var badgerOpts badger.Options
	if inMem {
		badgerOpts = badger.DefaultOptions("").WithInMemory(inMem)
	} else {
		badgerOpts = badger.DefaultOptions(path)
	}

	if logging {
		badgerOpts = badgerOpts.WithLogger(logger)
	}

	db, err := badger.Open(badgerOpts)
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping badgerDB: %w", err)
	}
	return db, nil
}

func initSQLite() (*sql.DB, error) {
	// Create the directory if it does not exist
	dbPath := "data/db"
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err := os.MkdirAll(dbPath, 0755) // Create the directory with appropriate permissions
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("sqlite3", "./data/db/av.db")
	if err != nil {
		return nil, err
	}

	// Use a ping to ensure the database is accessible
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = migrations(db)
	if err != nil {
		return nil, err
	}

	return db, err
}

func migrations(db *sql.DB) error {
	createTableSQL := `CREATE TABLE IF NOT EXISTS zipped (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    zipFileName TEXT NOT NULL,
    fileName TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		return err
	}
	return nil
}

type store struct {
	badgerDB    *badger.DB
	badgerMutex *sync.Mutex

	sqlDB    *sql.DB
	sqlMutex *sync.Mutex
}

type Store interface {
	// Get , Put and Delete for BadgerDb
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error

	// Get , Put and Delete for SQLite
	InsertZipRecord(zipFileName, fileName string) error
	DeleteZipRecord(zipFileName string) error
}
