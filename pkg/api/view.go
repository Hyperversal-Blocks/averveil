package api

import (
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	view2 "github.com/hyperversal-blocks/averveil/pkg/view"
)

func NewViewController(logger *logrus.Logger, service view2.Service) View {
	return &view{
		logger:  logger,
		service: service,
	}
}

type view struct {
	logger  *logrus.Logger
	service view2.Service
}

func (v *view) CSV(w http.ResponseWriter, r *http.Request) {
	fileName := r.FormValue("key")

	if len(strings.TrimSpace(fileName)) == 0 {
		v.logger.Error("query param key cannot be empty")
		WriteJson(w, "query param key cannot be empty", http.StatusBadRequest)
		return
	}

	decodedMap, err := v.service.CSV(fileName)
	if err != nil {
		WriteJson(w, outputDTO{
			Message:   err.Error(),
			Data:      nil,
			Timestamp: time.RFC3339,
		}, http.StatusOK)
		return
	}

	WriteJson(w, outputDTO{
		Message:   "success",
		Data:      decodedMap,
		Timestamp: time.RFC3339,
	}, http.StatusOK)
}

type View interface {
	CSV(w http.ResponseWriter, r *http.Request)
}
