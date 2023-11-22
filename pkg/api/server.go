package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/hyperversal-blocks/averveil/pkg/auth"
	"github.com/hyperversal-blocks/averveil/pkg/configuration"
	"github.com/hyperversal-blocks/averveil/pkg/hblock"
	jwtPkg "github.com/hyperversal-blocks/averveil/pkg/jwt"
	"github.com/hyperversal-blocks/averveil/pkg/logger"
	"github.com/hyperversal-blocks/averveil/pkg/node"
	"github.com/hyperversal-blocks/averveil/pkg/store"
	swarmService "github.com/hyperversal-blocks/averveil/pkg/swarm"
	upl "github.com/hyperversal-blocks/averveil/pkg/upload"
	u "github.com/hyperversal-blocks/averveil/pkg/user"
	"github.com/hyperversal-blocks/averveil/pkg/util"
	v "github.com/hyperversal-blocks/averveil/pkg/view"
)

type Services struct {
	Config configuration.Config
	logger *logrus.Logger
	Api    *Api
}

func Init() error {
	services, err := bootstrapper(context.Background())
	if err != nil {
		return err
	}

	services.Api.Cors()
	services.Api.Routes()

	go func() {
		services.startServer()
	}()
	select {}
}

func (c *Services) startServer() {
	serverConf := c.Config.GetConfig().Server
	address := serverConf.Host + serverConf.PORT

	c.logger.Info("Starting Server at:", address)

	err := http.ListenAndServe(address, c.Api.GetRouter())
	if err != nil {
		c.logger.Error("error starting server at ", address, " with error: ", err)
		panic(err)
	}
}

func bootstrapper(ctx context.Context) (*Services, error) {
	conf, err := configuration.Init()
	if err != nil {
		return nil, fmt.Errorf("error bootstrapping config: %w", err)
	}

	confInstance := conf.GetConfig()

	loggerInstance := logger.Init(confInstance.Logger.Level, confInstance.Logger.Env)
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
	userService := u.New(storer, loggerInstance, node.Signer.EthereumAddress())

	authService := auth.New(node.Signer, storer, jwt, userService)

	address, abi, err := util.ContractParser(confInstance.Contracts.HBLOCK.Address, confInstance.Contracts.HBLOCK.ABI)
	if err != nil {
		return nil, fmt.Errorf("unable to parse the HBLOCK contract: %w", err)
	}

	hblockContractService := hblock.New(&node.TxService, node.Signer.EthereumAddress(), loggerInstance, address, abi)

	uploadService := upl.NewUploadService(loggerInstance, storer)

	viewService := v.NewViewService(loggerInstance, storer)

	swarmService := swarmService.New(confInstance.Swarm.Host + confInstance.Swarm.PORT)

	// Bootstrapping Controllers
	userController := NewUserController(loggerInstance, hblockContractService, confInstance)

	uploadController := NewUploadController(loggerInstance, uploadService)

	viewController := NewViewController(loggerInstance, viewService)

	swarmController := NewSwarmController(loggerInstance, swarmService)

	apiService := InitAPI(loggerInstance,
		authService,
		userController,
		uploadController,
		viewController,
		jwt,
		node, swarmController)

	return &Services{
		Config: confInstance,
		logger: loggerInstance,
		Api:    apiService,
	}, nil
}

func InitAPI(
	loggerInstance *logrus.Logger,
	authService auth.Auth,
	userController User,
	uploadController Upload,
	viewController View,
	jwt jwtPkg.JWT,
	node *node.Node,
	swarmController Swarm,
) *Api {
	return New(loggerInstance, chi.NewMux(), authService, userController, node, jwt, uploadController, viewController, swarmController)
}
