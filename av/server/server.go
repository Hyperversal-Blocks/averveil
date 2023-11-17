package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/configuration"
	"github.com/hyperversal-blocks/averveil/pkg/api"
	"github.com/hyperversal-blocks/averveil/pkg/auth"
	"github.com/hyperversal-blocks/averveil/pkg/hblock"
	jwtPkg "github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/logger"
	"github.com/hyperversal-blocks/averveil/pkg/node"
	"github.com/hyperversal-blocks/averveil/pkg/store"
	"github.com/hyperversal-blocks/averveil/pkg/upload"
	"github.com/hyperversal-blocks/averveil/pkg/user"
	"github.com/hyperversal-blocks/averveil/pkg/util"
	"github.com/hyperversal-blocks/averveil/pkg/view"
)

type Services struct {
	config *configuration.Config
	logger *logrus.Logger
	api    *api.Services
}

func Init(desktopConfig bool) error {
	if desktopConfig {
		// TODO: Add logic for when config will be set from frontend
		// TODO: move server init to pkg
	}

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

	storer, err := store.New(loggerInstance, confInstance.Store.Path, confInstance.Store.InMem, confInstance.Store.Logging)
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

	// Bootstrapping Services
	userService := user.New(storer, loggerInstance, node.Signer.EthereumAddress())

	authService := auth.New(node.Signer, storer, jwt, userService)

	address, abi, err := util.ContractParser(confInstance.Contracts.HBLOCK.Address, confInstance.Contracts.HBLOCK.ABI)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the HBLOCK contract: %w", err)
	}

	hblockContractService := hblock.New(&node.TxService, node.Signer.EthereumAddress(), loggerInstance, address, abi)

	uploadService := upload.NewUploadService(loggerInstance, storer)

	viewService := view.NewViewService(loggerInstance, storer)
	// Bootstrapping Controllers
	userController := api.NewUserController(loggerInstance, hblockContractService)

	uploadController := api.NewUploadController(loggerInstance, uploadService)

	viewController := api.NewViewController(loggerInstance, viewService)

	apiService := api.New(loggerInstance, chi.NewMux(), authService, userController, node, jwt, uploadController, viewController)

	return &Services{
		config: confInstance,
		logger: loggerInstance,
		api:    apiService,
	}, nil
}
