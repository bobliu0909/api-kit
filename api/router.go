package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/rl5c/api-server/api/base"
	_ "github.com/rl5c/api-server/api/v1"
	"github.com/rl5c/api-server/conf"
	"github.com/rl5c/api-server/pkg/controllers"
)

type IRouter interface {
	//ServeHTTP used to handle the http requests
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Router struct {
	engine   *gin.Engine
	handlers map[string]interface{}
}

func NewRouter(cluster string, controller controllers.BaseController, config *conf.APIConfig) IRouter {
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
			handlers[version] = constructor(controller)
		}
	}
	router := &Router{
		engine:   engine,
		handlers: handlers,
	}
	router.initRouter(cluster)
	return router
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.engine.ServeHTTP(w, r)
}

func (router *Router) initRouter(cluster string) {
	for version, handler := range router.handlers {
		if handler != nil {
			prefix := fmt.Sprintf("/%s/%s", cluster, version)
			crdsGroup := router.engine.Group(fmt.Sprintf("%s/crds", prefix))
			crsGroup := router.engine.Group(fmt.Sprintf("%s/crs", prefix))
			handler.(base.Initializer).SetRouter(crdsGroup, crsGroup)
		}
	}
}
