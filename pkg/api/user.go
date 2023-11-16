package api

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/hblock"
)

func NewUserController(logger *logrus.Logger, hblock hblock.Contract) User {
	return &user{
		logger: logger,
		hblock: hblock,
	}
}

type User interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
}

type user struct {
	logger *logrus.Logger
	hblock hblock.Contract
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
