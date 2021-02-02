package app

import (
	"context"

	"github.com/rl5c/api-server/api"
	"github.com/rl5c/api-server/conf"
	"github.com/rl5c/api-server/pkg/logger"
	"github.com/rl5c/api-server/pkg/controllers"
	"github.com/rl5c/api-server/pkg/service"
)

type AppDaemon struct {
	apiServer	*api.APIServer
	controller  controllers.BaseController
	stopCh      <-chan struct{}
}

func New(stopCh <-chan struct{}) (*AppDaemon, error) {
	manageSvc, err := service.NewManageService(stopCh)
	if err != nil {
		return nil, err
	}
	simpleSvc, err := service.NewSimpleService(stopCh)
	if err != nil {
		return nil, err
	}
	controller := controllers.NewController(manageSvc, simpleSvc)
	apiServer, err := api.NewServer(context.Background(), conf.Cluster(), controller, conf.APIConfigValue())
	if err != nil {
		return nil, err
	}
	return &AppDaemon{
		apiServer:	apiServer,
		controller: controller,
		stopCh:		stopCh,
	}, nil
}

func (daemon *AppDaemon) Startup() error {
	if err := daemon.controller.Manage().Open(context.Background()); err != nil {
		return err
	}
	return daemon.apiServer.Startup()
}

func (daemon *AppDaemon) Stop() {
	daemon.controller.Manage().Close()
	if err := daemon.apiServer.Stop(); err != nil {
		logger.ERROR("[#app#] api server closing error, %s.", err.Error())
	}
}
