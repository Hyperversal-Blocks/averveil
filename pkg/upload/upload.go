package upload

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/store"
)

var (
	ErrFileEmpty = fmt.Errorf("file is empty")
)

type upload struct {
	store  store.Store
	logger *logrus.Logger
}

func (u *upload) CSV(csvFile *FileCSV) error {
	// TODO: add logic to overwrite existing data if true, append to existing data or create new copy

	reader := csv.NewReader(csvFile.File)

	// Process each row
	keyValuePairs := make(map[string]string)
	fileIsEmpty := true
	for {
		row, err := reader.Read()
		if err != nil {
			break // End of file or an error
		}

		// Check if row is empty or contains only spaces
		isEmptyRow := true
		for _, field := range row {
			if strings.TrimSpace(field) != "" {
				isEmptyRow = false
				break
			}
		}

		if !isEmptyRow {
			fileIsEmpty = false
			key := row[0]                       // First column as the key
			value := strings.Join(row[1:], ",") // Joining all other columns as the value
			keyValuePairs[key] = value
		}
	}

	// Check if the file was empty
	if fileIsEmpty {
		return ErrFileEmpty
	}

	// Convert map to JSON bytes
	data, err := json.Marshal(keyValuePairs)
	if err != nil {
		return err
	}

	err = u.store.Put(csvFile.FileName, data)
	if err != nil {
		return err
	}

	return nil
}

type FileCSV struct {
	FileName string
	Ext      string
	File     multipart.File
}

func NewUploadService(logger *logrus.Logger, store store.Store) Service {
	return &upload{
		logger: logger,
		store:  store,
	}
}

type Service interface {
	CSV(csv *FileCSV) error
}
