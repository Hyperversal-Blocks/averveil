package user

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

type profile struct {
	db      *badger.DB
	logger  *logrus.Logger
	address common.Address
}

func (p *profile) Create(ctx context.Context, user *User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("unable to marshal user object: %w", err)
	}

	err = p.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(p.address.String()), userJSON)
	})
	if err != nil {
		return fmt.Errorf("unable to create user object: %w", err)
	}

	return nil
}

func (p *profile) Update(ctx context.Context, update *User) error {
	// Start a writable transaction.
	err := p.db.Update(func(txn *badger.Txn) error {
		// Get the existing User data
		item, err := txn.Get([]byte(p.address.String()))
		if err != nil {
			return err
		}
		userJSON, err := item.ValueCopy(nil)
		if err != nil {
			return err
		}

		// Unmarshal the JSON data into a User struct
		var user User
		err = json.Unmarshal(userJSON, &user)
		if err != nil {
			return err
		}

		user = *update

		// Marshal the modified User struct back to JSON
		updatedUserJSON, err := json.Marshal(user)
		if err != nil {
			return err
		}

		// Set the updated JSON data back to the same key
		return txn.Set([]byte(p.address.String()), updatedUserJSON)
	})

	if err != nil {
		return fmt.Errorf("unable to update user object: %w", err)
	}

	return nil
}

func (p *profile) Get(ctx context.Context) (*User, error) {
	// Retrieve the User data from BadgerDB
	var userJSON []byte
	err := p.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(p.address.String()))
		if err != nil {
			return err
		}
		userJSON, err = item.ValueCopy(nil)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("unable to fetch user object: %w", err)
	}

	// Unmarshal the JSON data back into a User struct
	var user User
	err = json.Unmarshal(userJSON, &user)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal user object: %w", err)
	}
	return &user, nil
}

func New(db *badger.DB, logger *logrus.Logger, address common.Address) Service {
	return &profile{
		db:      db,
		logger:  logger,
		address: address,
	}
}

type User struct {
	Name    string  `json:"name,omitempty"`
	Age     int     `json:"age,omitempty"`
	Gender  string  `json:"gender,omitempty"`
	DOB     string  `json:"dob"`
	Address Address `json:"address"`
	Contact Contact `json:"contact"`
	Type    Type    `json:"type"`
	Wallet  string  `json:"wallet"`
	PubKey  string  `json:"pubKey"`
}

type Address struct {
	City       string `json:"city,omitempty"`
	Country    string `json:"country,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}

type Contact struct {
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Type struct {
	MBTI      string `json:"mbti,omitempty"`
	Horoscope string `json:"horoscope,omitempty"`
}

type Service interface {
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Get(ctx context.Context) (*User, error)
}
