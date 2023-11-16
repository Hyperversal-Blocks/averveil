package api

import (
	"encoding/csv"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

func NewUploadController(logger *logrus.Logger, store store.Store) Upload {
	return &upload{
		logger: logger,
		store:  store,
	}
}

type upload struct {
	logger *logrus.Logger
	store  store.Store
}

func (u *upload) CSV(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		u.logger.Info("unable to parse form data: ", err)

		WriteJson(w, "unable to parse form data", http.StatusInternalServerError)
		return
	}

	fileName := r.FormValue("name")

	reader := csv.NewReader(file)

	// Process each row
	keyValuePairs := make(map[string]string)
	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or an error
		}

		key := row[0]                       // First column as the key
		value := strings.Join(row[1:], ",") // Joining all other columns as the value

		keyValuePairs[key] = value
	}

	// Convert map to JSON bytes
	data, err := json.Marshal(keyValuePairs)
	if err != nil {
		u.logger.Error("internal server error: ", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}

	err = u.store.Put(fileName, data)
	if err != nil {
		u.logger.Error("internal server error: ", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}

	WriteJson(w, "uploaded successfully", http.StatusOK)
}

type Upload interface {
	CSV(w http.ResponseWriter, r *http.Request)
}
