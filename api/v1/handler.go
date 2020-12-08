package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-gin/api/base"
	"github.com/rl5c/api-gin/pkg/cluster"
)

type V1Handler struct {
	base.Handler
}

func NewDefaultHandler(clusterService cluster.IClusterService) base.HandlerInterface {
	return &V1Handler{
		Handler: base.Handler{
			ClusterService: clusterService,
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
