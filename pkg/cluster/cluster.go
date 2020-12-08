package cluster

const (
	NGClusterService = "ngCluster"
)

type IClusterService interface {
	Startup() error
	Close()
}