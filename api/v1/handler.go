package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *V1Handler) RegisterHandlerFunc(ctx *gin.Context) {
	err := handler.Controller.Simple().Register(ctx, "aaa")
	ctx.JSON(http.StatusOK, err)
}

func (handler *V1Handler) UnRegisterHandlerFunc(ctx *gin.Context) {
	err := handler.Controller.Simple().UnRegister(ctx, "aaa")
	ctx.JSON(http.StatusOK, err)
}

/*
func (handler *V1Handler) CRDListHandlerFunc(ctx *gin.Context) {
	crds := handler.Service.CRDList(ctx)
	ctx.JSON(http.StatusOK, crds)
}

func (handler *V1Handler) CRDHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	crd := handler.Service.GetCRD(ctx, resource)
	if crd == nil {
		ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, nil))
		return
	}
	ctx.JSON(http.StatusOK, crd)
}

func (handler *V1Handler) CreateCRDHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(ErrClientRequestResolveCode, RequestBodyInvalidMsg, err))
		return
	}
	crdDispose, err := handler.Service.CreateCRD(ctx, resource, data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, DataResponse(ResponseSuccessfullyCode, fmt.Sprintf("%s install successfully", resource), crdDispose))
}

func (handler *V1Handler) RemoveCRDHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	crdDispose, err := handler.Service.DeleteCRD(ctx, resource)
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, DataResponse(ResponseSuccessfullyCode, fmt.Sprintf("%s uninstall successfully", resource), crdDispose))
}

func (handler *V1Handler) CRListHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	namespace := ctx.Param("namespace")
	crs, err := handler.Service.CRList(ctx, resource, namespace)
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, crs)
}

func (handler *V1Handler) CRHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	if name != "history" {
		cr, err := handler.Service.GetCR(ctx, resource, namespace, name)
		if err != nil {
			if strings.Index(err.Error(), "not found") != -1 {
				ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
			return
		}
		ctx.JSON(http.StatusOK, cr)
	} else {
		jobs, err := handler.Service.HistoryCRs(ctx, resource, namespace)
		if err != nil {
			if strings.Index(err.Error(), "not found") != -1 {
				ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
				return
			}
			ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
			return
		}
		ctx.JSON(http.StatusOK, jobs)
	}
}

func (handler *V1Handler) CreateCRHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(ErrClientRequestResolveCode, RequestBodyInvalidMsg, err))
		return
	}
	cr, err := handler.Service.CreateCR(ctx, resource, data)
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, cr)
}

func (handler *V1Handler) ReshapeCRsHandlerFunc(ctx *gin.Context) {
	var request types.ReshapeOptions
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(ErrClientRequestResolveCode, RequestBodyInvalidMsg, err))
		return
	}
	resource := ctx.Param("resource")
	if err := handler.Service.ReshapeCRs(ctx, resource, request); err != nil {
		if err == types.ErrNoValidResourceOperated {
			ctx.JSON(http.StatusNotAcceptable, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func (handler *V1Handler) UpdateCRHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	data, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(ErrClientRequestResolveCode, RequestBodyInvalidMsg, err))
		return
	}
	cr, err := handler.Service.UpdateCR(ctx, resource, data)
	if err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, cr)
}

func (handler *V1Handler) DeleteCRHandlerFunc(ctx *gin.Context) {
	resource := ctx.Param("resource")
	namespace := ctx.Param("namespace")
	name := ctx.Param("name")
	if err := handler.Service.DeleteCR(ctx, resource, namespace, name); err != nil {
		if strings.Index(err.Error(), "not found") != -1 {
			ctx.JSON(http.StatusNotFound, ErrorResponse(ErrServerResourceNotFoundCode, ResourceNotFoundMsg, err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(ErrServerInternalFailedCode, ServiceInternalErrorMsg, err))
		return
	}
	ctx.JSON(http.StatusOK, DataResponse(ResponseSuccessfullyCode, fmt.Sprintf("%s remove successfully", name), nil))
}
*/
