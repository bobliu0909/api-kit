package base

import (
	"github.com/gin-gonic/gin"

	"github.com/rl5c/api-gin/pkg/controllers"
)

type Initializer interface {
	SetRouter(group *gin.RouterGroup)
}

type Handler struct {
	Initializer
	Controller controllers.IController
}

type HandlerInterface interface{}

type ConstructorHandlerFunc func(controller controllers.IController) HandlerInterface

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
