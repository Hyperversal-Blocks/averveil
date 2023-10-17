package user

import (
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

type profile struct {
	db      store.Store
	logger  *logrus.Logger
	address common.Address
}

func (p *profile) Create(user *User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("unable to marshal user object: %w", err)
	}

	err = p.db.Put(p.address.String(), userJSON)
	if err != nil {
		return fmt.Errorf("unable to create user object: %w", err)
	}

	return nil
}

func (p *profile) Get() (*User, error) {
	// Retrieve the User data from BadgerDB
	obj, err := p.db.Get(p.address.String())
	if err != nil {
		return nil, fmt.Errorf("unable to fetch user object: %w", err)
	}

	// Unmarshal the JSON data back into a User struct
	var user User
	err = json.Unmarshal(obj, &user)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal user object: %w", err)
	}
	return &user, nil
}

func New(db store.Store, logger *logrus.Logger, address common.Address) Service {
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
	Create(user *User) error
	Get() (*User, error)
}
