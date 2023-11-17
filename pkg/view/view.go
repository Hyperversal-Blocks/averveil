package view

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

type view struct {
	store  store.Store
	logger *logrus.Logger
}

func (v *view) CSV(fileName string) (map[string]string, error) {
	data, err := v.store.Get(fileName)
	if err != nil {
		// TODO: add cases for key not found
		return nil, err
	}

	// Convert JSON bytes back to map
	var decodedMap map[string]string

	err = json.Unmarshal(data, &decodedMap)
	if err != nil {
		return nil, err
	}
	return decodedMap, nil
}

type Service interface {
	CSV(fileName string) (map[string]string, error)
}

func NewViewService(logger *logrus.Logger, store store.Store) Service {
	return &view{
		logger: logger,
		store:  store,
	}
}
