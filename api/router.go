package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-gin/api/base"
	_ "github.com/rl5c/api-gin/api/v1"
	_ "github.com/rl5c/api-gin/api/v2"
	"github.com/rl5c/api-gin/conf"
	"github.com/rl5c/api-gin/pkg/cluster"
)

type IRouter interface {
	//ServeHTTP used to handle the http requests
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	engine   *gin.Engine
	handlers map[string]interface{}
}

func NewRouter(clusterService cluster.IClusterService, config *conf.API) IRouter {
	mode := gin.DebugMode
	if !config.Debug {
		mode = gin.ReleaseMode
	}
	gin.SetMode(mode)
	engine := gin.Default()
	for _, key := range config.Middleware {
		middleware, ret := middlewares[strings.ToLower(key)]
		if ret {
			engine.Use(middleware())
		}
	}
	handlers := map[string]interface{}{}
	for _, version := range config.Version {
		constructor := base.HandlerConstructor(version)
		if constructor != nil {
			handlers[version] = constructor(clusterService)
		}
	}
	router := &Router{
		engine:   engine,
		handlers: handlers,
	}
	router.initRouter()
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.engine.ServeHTTP(w, r)
}

func (router *Router) initRouter() {
	for version, handler := range router.handlers {
		if handler != nil {
			group := router.engine.Group("/" + version)
			handler.(base.Initializer).SetRouter(group)
		}
	}
}
