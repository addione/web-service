package main

import (
	"github.com/addione/web-service/config"
	"github.com/gin-gonic/gin"
)

func handleHttp() {
	server := gin.Default()
	config.AddRoutes(server)
	server.Run(":8091")

}
