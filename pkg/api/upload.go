package api

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	upload2 "github.com/hyperversal-blocks/averveil/pkg/upload"
)

func NewUploadController(logger *logrus.Logger, service upload2.Service) Upload {
	return &upload{
		logger:  logger,
		service: service,
	}
}

type upload struct {
	logger  *logrus.Logger
	service upload2.Service
}

func (u *upload) ZIP(w http.ResponseWriter, r *http.Request) {
	// TODO implement me
	panic("implement me")
}

func (u *upload) CSV(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		u.logger.Info("unable to parse form data: ", err)

		WriteJson(w, "unable to parse form data", http.StatusInternalServerError)
		return
	}

	err = u.service.CSV(&upload2.FileCSV{
		FileName: header.Filename,
		Ext:      header.Header.Get("Content-Type"),
		File:     file,
	})

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
		Data:      nil,
		Timestamp: time.RFC3339,
	}, http.StatusOK)
}

type Upload interface {
	CSV(w http.ResponseWriter, r *http.Request)
	ZIP(w http.ResponseWriter, r *http.Request)
}
