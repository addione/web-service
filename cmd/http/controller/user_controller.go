package controller

import (
	"net/http"

	"github.com/addione/web-service/cmd/http/manager"
	"github.com/addione/web-service/model/request"
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
	createUserParams, err := request.ValidateCreateUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	u, err := uc.um.CreateNewUser(createUserParams)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, u)

}
