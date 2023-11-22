package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"

	"github.com/hyperversal-blocks/averveil/pkg/auth"
	jwtPkg "github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/node"
)

type Api struct {
	logger *logrus.Logger
	router *chi.Mux
	auth   auth.Auth
	upload Upload
	view   View
	user   User
	node   *node.Node
	jwt    jwtPkg.JWT
	swarm  Swarm

	hblockSemaphore *semaphore.Weighted
}

func New(logger *logrus.Logger, router *chi.Mux, auth auth.Auth, user User, node *node.Node, jwtService jwtPkg.JWT, data Upload, view View, swarm Swarm) *Api {
	return &Api{
		logger:          logger,
		router:          router,
		auth:            auth,
		node:            node,
		jwt:             jwtService,
		hblockSemaphore: semaphore.NewWeighted(1),
		user:            user,
		upload:          data,
		view:            view,
		swarm:           swarm,
	}
}
