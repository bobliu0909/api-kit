package app

import (
	"context"

	"github.com/rl5c/api-gin/api"
	"github.com/rl5c/api-gin/conf"
	"github.com/rl5c/api-gin/logger"
	"github.com/rl5c/api-gin/pkg/cluster"
	"github.com/rl5c/api-gin/pkg/cluster/ngcluster"
	"github.com/rl5c/api-gin/pkg/storage"
	"github.com/rl5c/api-gin/pkg/storage/driver"
)

type AppDaemon struct {
	stopCh    chan struct{}
	storageDriver driver.IStorageDriver
	clusterService cluster.IClusterService
	apiServer *api.APIServer
}

func New() (*AppDaemon, error) {
	stopCh := make(chan struct{})
	storageDriver, err := storage.StorageFactory(conf.Location(), conf.StorageConfig())
	if err != nil {
		return nil, err
	}
	clusterService, err := ngcluster.NewClusterService(storageDriver, stopCh)
	if err != nil {
		return nil, err
	}
	apiServer, err := api.NewServer(context.Background(), clusterService, conf.APIConfig())
	if err != nil {
		return nil, err
	}
	return &AppDaemon{
		stopCh: stopCh,
		storageDriver: storageDriver,
		clusterService: clusterService,
		apiServer: apiServer,
	}, nil
}

func (daemon *AppDaemon) Startup() error {
	var err error
	defer func() {
		if err != nil {
			daemon.storageDriver.Close()
			daemon.clusterService.Close()
		}
	}()

	if err = daemon.storageDriver.Open(); err != nil {
		return err
	}

	if err = daemon.clusterService.Startup(); err != nil {
		return err
	}
	return daemon.apiServer.Startup()
}

func (daemon *AppDaemon) Stop() {
	close(daemon.stopCh)
	if err := daemon.apiServer.Stop(); err != nil {
		logger.ERROR("[#app#] api server close error, %s.", err.Error())
	}
	daemon.clusterService.Close()
	daemon.storageDriver.Close()
}
