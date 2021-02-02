package controllers

import (
	"context"

	"github.com/rl5c/api-server/pkg/service"
)

type SimpleController interface{
	Register(ctx context.Context, node string) error
	UnRegister(ctx context.Context, node string) error
	Release(node string)
}

type simpleController struct {
	manageSvc service.ManageService
	simpleSvc service.SimpleService
}

func NewSimpleController(manageService service.ManageService, simpleService service.SimpleService) SimpleController {
	return &simpleController{
		manageSvc: manageService,
		simpleSvc: simpleService,
	}
}

func (ctrl *simpleController) Register(ctx context.Context, node string) error {
	//todo first write db or cache, secondly operator service.
	return ctrl.simpleSvc.Register(ctx, node)
}

func (ctrl *simpleController) UnRegister(ctx context.Context, node string) error {
	//todo first write db or cache, secondly operator service.
	return ctrl.simpleSvc.UnRegister(ctx, node)
}

func (ctrl *simpleController) Release(node string) {
	ctrl.simpleSvc.Release(node)
}
