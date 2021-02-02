package controllers

import (
	"github.com/rl5c/api-server/pkg/service"
)

type BaseController interface {
	Manage() ManageController
	Simple() SimpleController
}

type baseController struct {
	manageController ManageController
	simpleController SimpleController
}

func NewController(manageService service.ManageService, simpleService service.SimpleService) BaseController {
	base := &baseController{
		manageController: NewManageController(manageService, simpleService),
		simpleController: NewSimpleController(manageService, simpleService),
	}
	return base
}

func (base *baseController) Manage() ManageController {
	return base.manageController
}

func (base *baseController) Simple() SimpleController {
	return base.simpleController
}
