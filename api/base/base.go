package base

import (
	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-server/pkg/controllers"
)

type Initializer interface {
	SetRouter(crdsGroup *gin.RouterGroup, crsGroup *gin.RouterGroup)
}

type Handler struct {
	Initializer
	Controller controllers.BaseController
}

type HandlerInterface interface{}

type ConstructorHandlerFunc func(Controller controllers.BaseController) HandlerInterface

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
