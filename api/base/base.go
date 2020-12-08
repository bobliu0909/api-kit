package base

import (
	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-gin/pkg/cluster"
)

type Initializer interface {
	SetRouter(group *gin.RouterGroup)
}

type Handler struct {
	Initializer
	ClusterService cluster.IClusterService
}

type HandlerInterface interface{}

type ConstructorHandlerFunc func(clusterService cluster.IClusterService) HandlerInterface

var (
	handlers = map[string]ConstructorHandlerFunc{}
)

func Register(version string, constructor ConstructorHandlerFunc) {
	handlers[version] = constructor
}

func HandlerConstructor(version string) ConstructorHandlerFunc {
	if constructor, ret := handlers[version]; ret {
		return constructor
	}
	return nil
}
