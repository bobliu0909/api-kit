package ngcluster

import (
	"github.com/rl5c/api-gin/pkg/cluster"
)

type ngCluster struct {
	stopCh <-chan struct{}
}

func NewCluster(stopCh <-chan struct{}) (cluster.ICluster, error) {
	return &ngCluster{
		stopCh: stopCh,
	}, nil
}

func (cluster *ngCluster) Startup() error {
	return nil
}

func (cluster *ngCluster) Close() {
}

