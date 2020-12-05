package controllers

import (
	"github.com/rl5c/api-gin/pkg/cluster"
	"github.com/rl5c/api-gin/pkg/storage/driver"
)

type IController interface {
	CreateService() error
	RemoveService() error
}

type basicController struct {
	cluster cluster.ICluster
	storage driver.IStorageDriver
}

func NewController(clusterEngine cluster.ICluster, storageDriver driver.IStorageDriver) IController {
	return &basicController{
		cluster: clusterEngine,
		storage: storageDriver,
	}
}

func (controller *basicController) CreateService() error {
	return nil
}

func (controller *basicController) RemoveService() error {
	return nil
}

