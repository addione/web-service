package controller

import (
	"net/http"

	"github.com/addione/web-service/cmd/http/manager"
	"github.com/gin-gonic/gin"
)

type userController struct {
	um manager.UserManager
}

func NewUserController(hdi *httpDIContainer) *userController {
	return &userController{
		um: *hdi.managerDIContainer.GetUserManager(),
	}
}

func (uc *userController) CreateUser(ctx *gin.Context) {
	uc.um.CreateNewUser()
	ctx.JSON(http.StatusOK, gin.H{`message`: `user creation is in progress`})
}
