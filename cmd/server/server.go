package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/averveil/configuration"
	"github.com/hyperversalblocks/averveil/pkg/api"
	auth2 "github.com/hyperversalblocks/averveil/pkg/auth"
	"github.com/hyperversalblocks/averveil/pkg/logger"
	"github.com/hyperversalblocks/averveil/pkg/node"
)

type Container struct {
	config     *configuration.Config
	logger     *logrus.Logger
	router     *chi.Mux
	node       *node.Node
	ethAddress common.Address
}

func Init() error {
	api.New()
	go func() {
		container.startServer()
	}()
	select {}
}

func (c *Container) startServer() {
	address := c.config.Server.Host + c.config.Server.PORT

	c.logger.Info("Starting Server at:", address)

	err := http.ListenAndServe(address, c.router)
	if err != nil {
		c.logger.Error("error starting server at ", address, " with error: ", err)
		panic(err)
	}
}

func (c *Container) routes() {
	c.router.Route("/storage", func(r chi.Router) {
		// setter
		// getter
	})
}

func bootstrapper(ctx context.Context) (
	*logrus.Logger,
	*configuration.Config,
	*node.Node,
	error) {
	confInstance, err := configuration.Init()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error bootstrapping config: %w", err)
	}

	loggerInstance := logger.Init(confInstance)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error bootstrapping logger: %w", err)
	}

	node, err := node.InitNode(ctx, *confInstance, loggerInstance)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("unable to initialize the node: %w", err)
	}

	auth, err := auth2.Auth()
	return loggerInstance, confInstance, node, nil
}

func initContainer(logger *logrus.Logger,
	config *configuration.Config,
	node *node.Node) *Container {
	return &Container{
		config:     config,
		logger:     logger,
		router:     chi.NewRouter(),
		node:       node,
		ethAddress: node.Signer.EthereumAddress(),
	}
}
