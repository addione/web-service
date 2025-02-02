package controller

import "github.com/addione/web-service/cmd/http/manager"

type httpDIContainer struct {
	controllers        Controllers
	managerDIContainer manager.ManagerDIContainer
}

type Controllers struct {
	userController *userController
}

func NewHttpDiContainer() *httpDIContainer {
	hdi := &httpDIContainer{managerDIContainer: *manager.NewManagerDIContainer()}
	hdi.controllers = NewContollersDI(hdi)
	return hdi
}

func NewContollersDI(hdi *httpDIContainer) Controllers {
	cdi := Controllers{
		userController: NewUserController(hdi),
	}
	return cdi
}

func (hdi *httpDIContainer) GetUserController() *userController {
	return hdi.controllers.userController
}
