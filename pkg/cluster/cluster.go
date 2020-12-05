package cluster

const (
	NGClusterEngine = "ngCluster"
)

type ICluster interface {
	Startup() error
	Close()
}