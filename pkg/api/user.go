package api

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/configuration"
	"github.com/hyperversal-blocks/averveil/pkg/hblock"
)

func NewUserController(logger *logrus.Logger, hblock hblock.Contract, conf configuration.Config) User {
	return &user{
		conf:   conf,
		logger: logger,
		hblock: hblock,
	}
}

type User interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
	GetConfig(w http.ResponseWriter, r *http.Request)
}

type user struct {
	logger *logrus.Logger
	hblock hblock.Contract
	conf   configuration.Config
}

func (u *user) GetConfig(w http.ResponseWriter, r *http.Request) {
	conf := u.conf.GetConfig()
	WriteJson(w, conf, http.StatusOK)
}

func (u *user) GetBalance(w http.ResponseWriter, r *http.Request) {
	balance, err := u.hblock.GetBalance(r.Context())
	if err != nil {
		u.logger.Error("internal server error: ", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}

	WriteJson(w, balance, http.StatusOK)
}
