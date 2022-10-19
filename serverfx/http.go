package serverfx

import (
	"github.com/gin-gonic/gin"
)

var _port = ":8080"

func NewHttpServer() *gin.Engine{
	gin.SetMode(gin.ReleaseMode)
	var serv = gin.Default()
	registerRouter(serv)
	return serv
}
