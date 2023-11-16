package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"

	"github.com/hyperversal-blocks/averveil/pkg/auth"
	jwtPkg "github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/node"
)

type Services struct {
	logger     *logrus.Logger
	router     *chi.Mux
	auth       auth.Auth
	upload     Upload
	view       View
	user       User
	node       *node.Node
	jwtService jwtPkg.JWT

	hblockSemaphore *semaphore.Weighted
}

func New(logger *logrus.Logger, router *chi.Mux, auth auth.Auth, user User, node *node.Node, jwtService jwtPkg.JWT, data Upload, view View) *Services {
	return &Services{
		logger:          logger,
		router:          router,
		auth:            auth,
		node:            node,
		jwtService:      jwtService,
		hblockSemaphore: semaphore.NewWeighted(1),
		user:            user,
		upload:          data,
		view:            view,
	}
}
