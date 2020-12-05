package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/rl5c/api-gin/api/base"
	"github.com/rl5c/api-gin/pkg/controllers"
)

type V1Handler struct {
	base.Handler
}

func NewDefaultHandler(controller controllers.IController) base.HandlerInterface {
	return &V1Handler{
		Handler: base.Handler{
			Controller: controller,
		},
	}
}

func (handler *V1Handler) PingHandleFunc(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "V1 PANG",
	})
}

func (handler *V1Handler) GetData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"aaa": "aaatest1",
	})
}
