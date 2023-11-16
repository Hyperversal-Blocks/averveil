package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

func NewViewController(logger *logrus.Logger, store store.Store) View {
	return &view{
		logger: logger,
		store:  store,
	}
}

type view struct {
	logger *logrus.Logger
	store  store.Store
}

func (u *view) CSV(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("key")

	if len(strings.TrimSpace(fileName)) == 0 {
		u.logger.Error("query param key cannot be empty")
		WriteJson(w, "query param key cannot be empty", http.StatusBadRequest)
		return
	}

	data, err := u.store.Get(fileName)
	if err != nil {
		// TODO: add cases for key not found
		u.logger.Error("internal server error: ", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}

	// Convert JSON bytes back to map
	var decodedMap map[string]string

	err = json.Unmarshal(data, &decodedMap)
	if err != nil {
		u.logger.Error("internal server error: ", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}

	WriteJson(w, decodedMap, http.StatusOK)
}

type View interface {
	CSV(w http.ResponseWriter, r *http.Request)
}
