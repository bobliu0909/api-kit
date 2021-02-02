package controllers

import (
	"context"

	"github.com/rl5c/api-server/pkg/service"
)

type ManageController interface{
	Open(ctx context.Context) error
	Close()
}

type manageController struct {
	manageSvc service.ManageService
	simpleSvc service.SimpleService
}

func NewManageController(manageService service.ManageService, simpleService service.SimpleService) ManageController {
	return &manageController{
		manageSvc: manageService,
		simpleSvc: simpleService,
	}
}

func (ctrl *manageController) Open(ctx context.Context) error {
	if err := ctrl.manageSvc.Open(ctx); err != nil {
		return err
	}
	//load all node add to simple service.
	return nil
}

func (ctrl *manageController) Close() {
	//clear and release node in simple service
	ctrl.manageSvc.Close()
}
