package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/averveil/configuration"
	"github.com/hyperversalblocks/averveil/pkg/jwt"
	"github.com/hyperversalblocks/averveil/pkg/logger"
	"github.com/hyperversalblocks/averveil/pkg/node"
)

type Container struct {
	config     *configuration.Config
	logger     *logrus.Logger
	router     *chi.Mux
	node       *node.Node
	ethAddress common.Address
	jwt        jwt.JWT
}

func Init() error {
	ctx := context.Background()

	logger, conf, node, err := bootstrapper(ctx)
	if err != nil {
		return fmt.Errorf("error bootstrapping core services: %w", err)
	}

	container := initContainer(logger, conf, node)
	container.cors()
	container.routes()

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

func (c *Container) cors() {
	c.router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))
}
