package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/rl5c/api-gin/api/base"
)

const (
	APIVersion = "v1"
)

func init() {
	base.Register(APIVersion, NewDefaultHandler)
}

func (handler *V1Handler) SetRouter(router *gin.RouterGroup) {
	router.GET("/ping", handler.PingHandleFunc)
	router.GET("/data", handler.GetData)
}