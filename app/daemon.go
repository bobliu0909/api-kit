package app

import (
	"context"

	"github.com/rl5c/api-gin/api"
	"github.com/rl5c/api-gin/conf"
	"github.com/rl5c/api-gin/logger"
	"github.com/rl5c/api-gin/pkg/cluster"
	"github.com/rl5c/api-gin/pkg/cluster/ngcluster"
	"github.com/rl5c/api-gin/pkg/controllers"
	"github.com/rl5c/api-gin/pkg/storage"
	"github.com/rl5c/api-gin/pkg/storage/driver"
)

type AppDaemon struct {
	stopCh    chan struct{}
	storageDriver driver.IStorageDriver
	clusterEngine cluster.ICluster
	apiServer *api.APIServer
}

func New() (*AppDaemon, error) {
	stopCh := make(chan struct{})
	storageDriver, err := storage.StorageFactory(conf.Location(), conf.StorageConfig())
	if err != nil {
		return nil, err
	}
	clusterEngine, err := ngcluster.NewCluster(stopCh)
	if err != nil {
		return nil, err
	}
	controller := controllers.NewController(clusterEngine, storageDriver)
	apiServer, err := api.NewServer(context.Background(), controller, conf.APIConfig())
	if err != nil {
		return nil, err
	}
	return &AppDaemon{
		stopCh: stopCh,
		storageDriver: storageDriver,
		clusterEngine: clusterEngine,
		apiServer: apiServer,
	}, nil
}

func (daemon *AppDaemon) Startup() error {
	defer func() {
		daemon.storageDriver.Close()
		daemon.clusterEngine.Close()
	}()
	if err := daemon.storageDriver.Open(); err != nil {
		return err
	}
	if err := daemon.clusterEngine.Startup(); err != nil {
		return err
	}
	go func() {
		err := daemon.apiServer.Startup()
		if err != nil {
			logger.ERROR("[#app#] api server startup error, %s.", err.Error())
		}
	}()
	return nil
}

func (daemon *AppDaemon) Stop() {
	close(daemon.stopCh)
	if err := daemon.apiServer.Stop(); err != nil {
		logger.ERROR("[#app#] api server close error, %s.", err.Error())
	}
	daemon.clusterEngine.Close()
	daemon.storageDriver.Close()
}
