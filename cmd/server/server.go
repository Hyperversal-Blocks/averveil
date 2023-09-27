package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/hyperversalblocks/averveil/configuration"
	"github.com/hyperversalblocks/averveil/pkg/api"
	"github.com/hyperversalblocks/averveil/pkg/auth"
	jwtPkg "github.com/hyperversalblocks/averveil/pkg/jwt"
	"github.com/hyperversalblocks/averveil/pkg/logger"
	"github.com/hyperversalblocks/averveil/pkg/node"
	"github.com/hyperversalblocks/averveil/pkg/store"
	"github.com/hyperversalblocks/averveil/pkg/user"
)

type Services struct {
	config *configuration.Config
	logger *logrus.Logger
	api    *api.Services
}

func Init() error {
	services, err := bootstrapper(context.Background())
	if err != nil {
		return err
	}

	services.api.Cors()
	services.api.Routes()

	go func() {
		services.startServer()
	}()
	select {}
}

func (c *Services) startServer() {
	address := c.config.Server.Host + c.config.Server.PORT

	c.logger.Info("Starting Server at:", address)

	err := http.ListenAndServe(address, c.api.GetRouter())
	if err != nil {
		c.logger.Error("error starting server at ", address, " with error: ", err)
		panic(err)
	}
}

func bootstrapper(ctx context.Context) (*Services, error) {
	confInstance, err := configuration.Init()
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping config: %w", err)
	}

	loggerInstance := logger.Init(confInstance)
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping logger: %w", err)
	}

	storer, err := store.New(ctx, loggerInstance, confInstance.Store.Path, confInstance.Store.InMem, confInstance.Store.Logging)
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping store: %w", err)
	}

	node, err := node.InitNode(ctx, confInstance.Chain.PrivateKey, confInstance.Chain.Endpoint, loggerInstance)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize the node: %w", err)
	}

	jwt := jwtPkg.New(confInstance.JWT.Issuer,
		confInstance.JWT.Issuer,
		confInstance.JWT.Expiry)

	userService := user.New(storer, loggerInstance, node.Signer.EthereumAddress())

	authService := auth.New(node.Signer, storer, jwt, userService)

	apiService := api.New(loggerInstance, chi.NewMux(), authService, userService, node, jwt)

	return &Services{
		config: confInstance,
		logger: loggerInstance,
		api:    apiService,
	}, nil
}
