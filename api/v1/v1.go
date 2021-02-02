package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rl5c/api-server/api/base"
	"github.com/rl5c/api-server/pkg/controllers"
)

const (
	APIVersion = "v1"
)

func init() {
	base.Register(APIVersion, NewDefaultHandler)
}

type V1Handler struct {
	base.Handler
}

func NewDefaultHandler(controller controllers.BaseController) base.HandlerInterface {
	return &V1Handler{
		Handler: base.Handler{
			Controller: controller,
		},
	}
}

func (handler *V1Handler) SetRouter(crdsGroup *gin.RouterGroup, crsGroup *gin.RouterGroup) {
	/*
	crdsGroup.GET("", handler.CRDListHandlerFunc)
	crdsGroup.GET(":resource", handler.CRDHandlerFunc)
	crdsGroup.POST(":resource", handler.CreateCRDHandlerFunc)
	crdsGroup.DELETE(":resource", handler.RemoveCRDHandlerFunc)
	crsGroup.GET(":resource/:namespace", handler.CRListHandlerFunc)
	crsGroup.GET(":resource/:namespace/:name", handler.CRHandlerFunc)
	crsGroup.POST(":resource/reshape", handler.ReshapeCRsHandlerFunc)
	crsGroup.POST(":resource", handler.CreateCRHandlerFunc)
	crsGroup.PUT(":resource", handler.UpdateCRHandlerFunc)
	crsGroup.DELETE(":resource/:namespace/:name", handler.DeleteCRHandlerFunc)
	*/
}
