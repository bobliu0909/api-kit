package v2

import (
	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-gin/api/base"
	"github.com/rl5c/api-gin/pkg/cluster"
)

type V2Handler struct {
	base.Handler
}

func NewDefaultHandler(clusterService cluster.IClusterService) base.HandlerInterface {
	return &V2Handler{
		Handler: base.Handler{
			ClusterService: clusterService,
		},
	}
}

func (handler *V2Handler) PingHandleFunc(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "V2 PANG",
	})
}

func (handler *V2Handler) GetData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"aaa": "aaatest2",
	})
}
