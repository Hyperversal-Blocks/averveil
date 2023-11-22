package api

import (
	"net/http"

	"github.com/sirupsen/logrus"

	swarmService "github.com/hyperversal-blocks/averveil/pkg/swarm"
)

func NewSwarmController(logger *logrus.Logger, swarmService swarmService.Swarm) Swarm {
	return &swarm{
		logger: logger,
		swarm:  swarmService,
	}
}

type swarm struct {
	logger *logrus.Logger
	swarm  swarmService.Swarm
}

func (s *swarm) GetNodeHealth(w http.ResponseWriter, r *http.Request) {
	resp, err := s.swarm.CheckNodeHealthAndReadiness()
	if err != nil {
		s.logger.Error("error fetching swarm health: %w", err)
		WriteJson(w, "internal server error", http.StatusInternalServerError)
		return
	}
	WriteJson(w, resp, http.StatusOK)
}

type Swarm interface {
	GetNodeHealth(w http.ResponseWriter, r *http.Request)
}
