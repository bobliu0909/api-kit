package ngcluster

import (
	"fmt"

	"github.com/rl5c/api-gin/pkg/cluster"
	"github.com/rl5c/api-gin/pkg/storage/driver"
)

type ngClusterService struct {
	storageDriver driver.IStorageDriver
	stopCh <-chan struct{}
}

func NewClusterService(storageDriver driver.IStorageDriver, stopCh <-chan struct{}) (cluster.IClusterService, error) {
	return &ngClusterService{
		storageDriver: storageDriver,
		stopCh: stopCh,
	}, nil
}

func (cluster *ngClusterService) Startup() error {
	fmt.Println("cluster started.")
	return nil
}

func (cluster *ngClusterService) Close() {
	fmt.Println("cluster closed.")
}
