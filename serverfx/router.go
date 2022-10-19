package serverfx

import "github.com/gin-gonic/gin"

func registerRouter(httpServ *gin.Engine){
	httpServ.GET("/ping", one.Default)
	httpServ.GET("/test", two.Default)
}
