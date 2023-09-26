package api

import (
	"crypto/ecdsa"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/averveil/pkg/auth"
	"github.com/hyperversalblocks/averveil/pkg/jwt"
	"github.com/hyperversalblocks/averveil/pkg/node"
	"github.com/hyperversalblocks/averveil/pkg/signer"
	"github.com/hyperversalblocks/averveil/pkg/user"
)

type Services struct {
	logger       *logrus.Logger
	router       *chi.Mux
	auth         auth.Auth
	user         user.Service
	signer       signer.Signer
	jwt          jwt.JWT
	node         node.Node
	pubKey       *ecdsa.PublicKey
	chainAddress common.Address
}

func New(logger *logrus.Logger, router *chi.Mux, auth auth.Auth, user user.Service, signer signer.Signer, jwt jwt.JWT, node node.Node) *Services {
	return &Services{
		logger:       logger,
		router:       router,
		auth:         auth,
		user:         user,
		signer:       signer,
		jwt:          jwt,
		node:         node,
		pubKey:       signer.GetPublicKey(),
		chainAddress: signer.EthereumAddress(),
	}
}

func (s *Services) Cors() {
	s.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}

func (s *Services) Routes() {
}
