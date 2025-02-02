package config

import (
	"github.com/addione/web-service/cmd/http/controller"
	"github.com/gin-gonic/gin"
)

func AddRoutes(server *gin.Engine) {
	// Add your routes here
	uc := controller.NewHttpDiContainer().GetUserController()
	server.POST("/user", uc.CreateUser)
}
